/*************************************************************************
    > File Name: kv_db.h
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Sun 20 Feb 2022 05:10:58 AM UTC
 ************************************************************************/

#ifndef _KV_DB_H
#define _KV_DB_H
#include <stdio.h>
#include <wiredtiger_ext.h>
#include <wiredtiger.h>
#include "dict.h"
#define SCHEMA_ENTRIES 0
#define SCHEMA_DOCS 1
#define SCHEMA_STATE 2

#define SCHEMA_LIMIT (1024)


typedef struct
{

  WT_SESSION *session;
  WT_CURSOR *cursor;
  void *ctx;
  char  *schema_format;
  char *schema_name[0]
} kv_schema_t;

typedef struct
{
  char *database_name;
  char *database_dir;
  dict_t *schema_ctx;
   WT_CONNECTION *conn;

  // struct kv_schema **schema_ctx;
} kv_db_t;



kv_schema_t *kv_schema_alloc(const char *name, void *ctx,bool is_force_drop);
void kv_schema_destroy(kv_schema_t *schema);

kv_db_t *kv_db_alloc(const char *database_name, const char *database_dir);
// schema register
kv_schema_t *kv_db_fetch_schema(kv_db_t *db, char *schema_name);
int kv_db_register_schema(kv_db_t *db, kv_schema_t *schema);
void kv_db_unregister_schema(kv_db_t *db, char *schema_name);
// key and value operation
int kv_db_set(kv_db_t *db, char *schema_name, void *key, size_t key_sz, void *val,size_t val_sz);
void *kv_db_get(kv_db_t *db, char *schema_name, void *key,size_t key_sz);
int kv_db_del(kv_db_t *db, char *schema_name, void *key,size_t key_sz);
void *kv_db_destroy(kv_db_t *db);
#endif
