/*************************************************************************
  > File Name: store.c
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: 二  3/ 8 11:11:03 2022
 ************************************************************************/

#include <stdio.h>
#include <assert.h>
#include <stdarg.h>
#include "../drpc/drpc.pb-c.h"
#include "../drpc/drpc_util.h"
#include "log.h"
#include "dbservice.h"
#include "kv_db.h"
#include "schmea_meta.h"
#include "dbservice.pb-c.h"
enum drpc_kv_method
{
  DRPC_METHOD_CREATE_SCHEMA = 201,
  DRPC_METHOD_DROP_SCHEMA = 202,

  DRPC_METHOD_PUT_KV = 301,
  DRPC_METHOD_GET_KV = 302,
  DRPC_METHOD_DEL_KV = 303,

};

static void dbservice_get_kv(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *ctx)
{
  kv_db_t *db = (kv_db_t *)ctx;
  struct drpc_alloc alloc = PROTO_ALLOCATOR_INIT(alloc);
  Dbservice__GetKvReq *req = dbservice__get_kv_req__unpack(&alloc.alloc, drpc_req->body.len, drpc_req->body.data);
  Dbservice__GetKvResp resp = DBSERVICE__GET_KV_RESP__INIT;
  log_info("*******************dbservice_get_kv**************");
  log_info("request schmea=%s,key=%s", req->schema_name, req->key);
  char msg_buf[2048] = {'\0'};
  schema_meta_rec_t *meta = dict_get(db->schmea_meta_cache, req->schema_name);
  if (meta == NULL)
  {
    resp.code = -1;
    snprintf(&msg_buf, 2048, "failed: schema %s not found", req->schema_name);
    resp.msg = &msg_buf;
  }
  else
  {
    size_t key_size = strlen(req->key);
    char *value = (char *)kv_db_get(db, req->schema_name, req->key, key_size);
    resp.schema_name = req->schema_name;
    resp.key = req->key;
    resp.value = value;
  }
  size_t resp_len = dbservice__get_kv_resp__get_packed_size(&resp);
  uint8_t *buf;
  D_ALLOC(buf, resp_len);
  if (buf != NULL)
  {
    dbservice__get_kv_resp__pack(&resp, buf);
    drpc_resp->body.len = resp_len;
    drpc_resp->body.data = buf;
  }
  dbservice__get_kv_req__free_unpacked(req, &alloc.alloc);
}

static void dbservice_put_kv(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *ctx)
{
  kv_db_t *db = (kv_db_t *)ctx;
  struct drpc_alloc alloc = PROTO_ALLOCATOR_INIT(alloc);
  Dbservice__PutKvReq *req = dbservice__put_kv_req__unpack(&alloc.alloc, drpc_req->body.len, drpc_req->body.data);
  Dbservice__PutKvResp resp = DBSERVICE__PUT_KV_RESP__INIT;
  log_info("request schmea=%s,key=%s,val=%s", req->schema_name, req->key, req->value);
  char msg_buf[2048] = {'\0'};
  schema_meta_rec_t *meta = dict_get(db->schmea_meta_cache, req->schema_name);
  if (meta == NULL)
  {
    resp.code = -1;
    snprintf(&msg_buf, 2048, "failed: schema %s not found", req->schema_name);
    resp.msg = &msg_buf;
  }
  else
  {
    size_t key_size = strlen(req->key);
    size_t value_size = strlen(req->value);
    if (kv_db_search(db, req->schema_name, req->key, key_size) != -1)
    {
      resp.code = -1;
      snprintf(&msg_buf, 2048, "failed: key %s  exists", req->key);
      resp.msg = &msg_buf;
    }
    else
    {
      size_t bytes = strlen(req->key) + strlen(req->value);
      schema_meta_rec_t *last_meta = schmea_meta_fetch(sys_schmea_meta_name, req->schema_name, db);
      __sync_fetch_and_add(&last_meta->kv_count, 1);
      __sync_fetch_and_add(&last_meta->bytes, bytes);
      schmea_meta_save(req->schema_name, req->key, last_meta, sizeof(*last_meta), db);
      kv_db_set(db,req->schema_name,req->key,key_size,req->value,value_size);
      resp.code = 0;
      snprintf(&msg_buf, 2048, "succ: put key %s  ok", req->key);
      resp.msg = &msg_buf;
    }
  }
  size_t resp_len = dbservice__put_kv_resp__get_packed_size(&resp);
  uint8_t *buf;
  D_ALLOC(buf, resp_len);
  if (buf != NULL)
  {
    dbservice__put_kv_resp__pack(&resp, buf);
    drpc_resp->body.len = resp_len;
    drpc_resp->body.data = buf;
  }
  dbservice__put_kv_req__free_unpacked(req, &alloc.alloc);
}

static void dbservice_create_schema(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *ctx)
{
  kv_db_t *db = (kv_db_t *)ctx;
  struct drpc_alloc alloc = PROTO_ALLOCATOR_INIT(alloc);
  Dbservice__CreateSchemaReq *req = dbservice__create_schema_req__unpack(&alloc.alloc, drpc_req->body.len, drpc_req->body.data);
  Dbservice__CreateSchemaResp resp = DBSERVICE__CREATE_SCHEMA_RESP__INIT;
  log_info("request schmea_name=%s", req->name);
  char msg_buf[2048] = {'\0'};
  schema_meta_rec_t *meta = dict_get(db->schmea_meta_cache, req->name);
  if (meta == NULL)
  {
    kv_schema_t *schema = kv_schema_alloc(req->name, db, false);
    assert(schema != NULL);
    schema_meta_rec_t *cur_meta = calloc(1, sizeof(*cur_meta));
    assert(cur_meta != NULL);
    schmea_meta_assign(cur_meta, 0, true, 0);
    schmea_cache_add(db->schmea_meta_cache, db, sys_schmea_meta_name, req->name, cur_meta, sizeof(*cur_meta));
    if (schema != NULL)
    {
      resp.code = 0;
      snprintf(&msg_buf, 2048, "succ: new schema %s  ok", req->name);

      resp.msg = &msg_buf;
    }
    else
    {
      resp.code = -1;
      snprintf(&msg_buf, 2048, "failed:init schema  %s ctx", req->name);
      resp.msg = &msg_buf;
    }
  }
  else
  {
    resp.code = -1;
    snprintf(&msg_buf, 2048, "failed: schema %s  exists", req->name);
    resp.msg = &msg_buf;
  }

  size_t resp_len = dbservice__create_schema_resp__get_packed_size(&resp);
  uint8_t *buf;
  D_ALLOC(buf, resp_len);
  if (buf != NULL)
  {
    dbservice__create_schema_resp__pack(&resp, buf);
    drpc_resp->body.len = resp_len;
    drpc_resp->body.data = buf;
  }
  dbservice__create_schema_req__free_unpacked(req, &alloc.alloc);
}

static void dbservice_drop_schema(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *ctx)
{
  kv_db_t *db = (kv_db_t *)ctx;
  struct drpc_alloc alloc = PROTO_ALLOCATOR_INIT(alloc);
  Dbservice__DropSchemaReq *req = dbservice__drop_schema_req__unpack(&alloc.alloc, drpc_req->body.len, drpc_req->body.data);
  Dbservice__DropSchemaResp resp = DBSERVICE__DROP_SCHEMA_RESP__INIT;
  log_info("drop schmea_name=%s", req->name);
  schema_meta_rec_t *meta = dict_get(db->schmea_meta_cache, req->name);
  char msg_buf[2048] = {'\0'};
  if (meta == NULL)
  {
    resp.code = -1;
    snprintf(&msg_buf, 2048, "failed: schema %s not found", req->name);
    resp.msg = &msg_buf;
  }
  else
  {
    kv_schema_t *schema = kv_schema_alloc(req->name, db, false);
    kv_schema_destroy(schema);
    schmea_cache_del(db->schmea_meta_cache, db, sys_schmea_meta_name, req->name);
    resp.code = 0;
    snprintf(&msg_buf, 2048, "succ: drop schema %s ok", req->name);
    resp.msg = &msg_buf;
    resp.name = req->name;
  }
  size_t resp_len = dbservice__drop_schema_resp__get_packed_size(&resp);
  uint8_t *buf;
  D_ALLOC(buf, resp_len);
  if (buf != NULL)
  {
    dbservice__drop_schema_resp__pack(&resp, buf);
    drpc_resp->body.len = resp_len;
    drpc_resp->body.data = buf;
  }
  dbservice__drop_schema_req__free_unpacked(req, &alloc.alloc);
}
void process_drpc_request(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *ctx)
{
  switch (drpc_req->method)
  {
  case DRPC_METHOD_CREATE_SCHEMA:
    dbservice_create_schema(drpc_req, drpc_resp, ctx);
    break;
  case DRPC_METHOD_DROP_SCHEMA:
    dbservice_drop_schema(drpc_req, drpc_resp, ctx);
  case DRPC_METHOD_PUT_KV:
    dbservice_put_kv(drpc_req, drpc_resp, ctx);
  case DRPC_METHOD_GET_KV:
    dbservice_get_kv(drpc_req, drpc_resp, ctx);
    break;
  default:
    break;
  }
}