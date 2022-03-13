/*************************************************************************
    > File Name: kv_db.c
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Sun 20 Feb 2022 05:15:21 AM UTC
 ************************************************************************/

#include <stdio.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <fcntl.h>
#include <assert.h>
#include <stdlib.h>
#include "kv_db.h"
#include "hashfn.h"
const char *schema_format = "key_format=u,value_format=u";

inline static void wt_item_init(WT_ITEM *item,void *data,size_t size) {
  item->data = data;
  item->size=size;
}

kv_schema_t *kv_schema_alloc(const char *schema_name, void *ctx,bool is_force_drop)
{
  assert(ctx != NULL);
  kv_db_t *db = (kv_db_t *)ctx;
  size_t schema_sz = strlen(schema_name);

  kv_schema_t *schema = (kv_schema_t *)calloc(1, sizeof(kv_schema_t) + schema_sz + 1);
  assert(schema != NULL);
 //assert(db->conn->reconfigure(db->conn,"session_max=40000")!=-1);
  assert(db->conn->open_session(db->conn, NULL, NULL, &schema->session) != -1);
  char schema_buf[256] = {'\0'};
  snprintf(&schema_buf, 256, "table:%s", schema_name);
  
  if(is_force_drop) {
      assert(schema->session->drop(schema->session, &schema_buf, schema_format) != -1);
  }
  assert(schema->session->create(schema->session, &schema_buf, schema_format) != -1);
  assert(schema->session->open_cursor(schema->session, &schema_buf, NULL, "overwrite=false", &schema->cursor) != -1);
  strncpy((char *)&schema->schema_name, schema_name, schema_sz);
  schema->schema_name[schema_sz] = '\0';
  schema->ctx = ctx;
  schema->schema_format = strdup(schema_format);
  fprintf(stdout, "create schema %s succ\n", (char *)&schema->schema_name);
  return schema;
}
void kv_schema_destroy(kv_schema_t *schema)
{
  if (schema != NULL)
  {
    kv_db_t *db = (kv_db_t *)schema->ctx;
    schema->cursor->close(schema->cursor);
    schema->session->drop(schema->session,&schema->schema_name,schema->schema_format);
    schema->session->close(schema->session,NULL);
    free(schema);
    schema = NULL;
  }
}

kv_db_t *kv_db_alloc(const char *database_name, const char *database_dir)
{
  if (database_name != NULL && database_dir != NULL)
  {

    char base_path[2048] = {'\0'};
    snprintf(&base_path, 2048, "%s/%s", database_dir, database_name);
    if (access(&base_path, F_OK) != 0 && mkdir(&base_path, 0755) != 0)
    {

      return -1;
    }

    kv_db_t *db = calloc(1, sizeof(kv_db_t));
    assert(db != NULL);
    assert(wiredtiger_open(&base_path, NULL, "create,salvage=false,session_max=40000,session_scratch_max=2MB", &db->conn) != -1);

    db->database_name = strdup(database_name);
    db->database_dir = strdup(database_dir);
    db->schema_ctx = dict_create(SCHEMA_LIMIT, hash_fnv1_64);
    return db;
  }
  return NULL;
}
int kv_db_set(kv_db_t *db, char *schema_name, void *key, size_t key_sz, void *val,size_t val_sz)
{

  kv_schema_t *schema = (kv_schema_t *)dict_get(db->schema_ctx, schema_name);
  WT_CURSOR *cursor = schema->cursor;
   WT_ITEM key_item, value_item;
   wt_item_init(&key_item,key,key_sz);
   wt_item_init(&value_item,val,val_sz);
  cursor->set_key(cursor, &key_item);
  cursor->set_value(cursor,&value_item);
  return cursor->insert(cursor);
}
void *kv_db_get(kv_db_t *db, char *schema_name, void *key,size_t key_sz)
{
  kv_schema_t *schema = (kv_schema_t *)dict_get(db->schema_ctx, schema_name);
  WT_CURSOR *cursor = schema->cursor;
  WT_ITEM key_item, value_item;
  wt_item_init(&key_item,key,key_sz);
  cursor->set_key(cursor,&key_item);
  if (cursor->search(cursor) != 0)
  {
    return NULL;
  }
  cursor->get_value(cursor, &value_item);
  return  value_item.data;
}
int kv_db_del(kv_db_t *db, char *schema_name, void *key,size_t key_sz)
{
  kv_schema_t *schema = (kv_schema_t *)dict_get(db->schema_ctx, schema_name);
  WT_CURSOR *cursor = schema->cursor;
    WT_ITEM key_item, value_item;
  wt_item_init(&key_item,key,key_sz);
  cursor->set_key(cursor, &key_item);
  return cursor->remove(cursor);
}
inline static void kv_schema_free_cb(void *ptr)
{
  kv_schema_t *schema = (kv_schema_t *)ptr;
  kv_schema_destroy(schema);
}
void *kv_db_destroy(kv_db_t *db)
{
  if (db != NULL)
  {

    dict_deinit(db->schema_ctx, kv_schema_free_cb);
    free(db->database_dir);
    free(db->database_name);
  }
}
kv_schema_t *kv_db_fetch_schema(kv_db_t *db, char *schema_name)
{
  if (db == NULL || schema_name == NULL)
  {
    return -1;
  }
  return (kv_schema_t *)dict_get(db->schema_ctx, schema_name);
}
int kv_db_register_schema(kv_db_t *db, kv_schema_t *schema)
{
  if (db == NULL || schema == NULL)
  {
    return -1;
  }
  char *schema_name = &schema->schema_name;
  if (dict_put(db->schema_ctx, schema_name, schema) != NULL)
  {
    return 0;
  }

  return -1;
}
void kv_db_unregister_schema(kv_db_t *db, char *schema_name)
{
  if (db != NULL && schema_name != NULL)
  {
    dict_del(db->schema_ctx, schema_name, NULL);
  }
}