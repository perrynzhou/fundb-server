/*************************************************************************
  > File Name: store.c
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: 二  3/ 8 11:11:03 2022
 ************************************************************************/

#include <stdio.h>
#include <assert.h>
#include "../drpc/drpc.pb-c.h"
#include "../drpc/drpc_util.h"
#include "log.h"
#include "kv.pb-c.h"
#include "store.h"
#include "kv_db.h"
#include "schmea_meta.h"
enum drpc_kv_method
{
  DRPC_METHOD_CREATE_SCHEMA = 201,
};

static void create_schema(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *ctx)
{
  kv_db_t *db = (kv_db_t *)ctx;
  struct drpc_alloc alloc = PROTO_ALLOCATOR_INIT(alloc);
  Kv__CreateSchemaReq *req = kv__create_schema_req__unpack(&alloc.alloc, drpc_req->body.len, drpc_req->body.data);
  Kv__CreateSchemaResp resp = KV__CREATE_SCHEMA_RESP__INIT;
  log_info("request schmea_name=%s", req->name);
  char msg_buf[2048] = {'\0'};
  // void *kv_db_get(kv_db_t *db, char *schema_name, void *key,size_t key_sz);
  size_t name_sz = strlen(req->name);
  schema_meta_t *meta = dict_get(db->schmea_meta_cache, req->name);
  if (meta == NULL)
  {
    kv_schema_t *schema = kv_schema_alloc(req->name, db, false);
    assert(schema != NULL);
    schema_meta_t cur_meta;
    schmea_meta_assign(&cur_meta, 0, true);
    schmea_cache_add(db->schmea_meta_cache, db, schmea_meta_name, req->name, &cur_meta, sizeof(cur_meta));
    if (schema != NULL)
    {
      resp.code = 0;
      resp.msg = "succ";
    }
    else
    {
      resp.code = -1;
      resp.msg = "fail:invalid schema ctx";
    }
  }
  else
  {
    resp.code = -1;
    resp.msg = "fail:schema exists";
  }
  size_t resp_len = kv__create_schema_resp__get_packed_size(&resp);
  uint8_t *buf;
  D_ALLOC(buf, resp_len);
  if (buf != NULL)
  {
    kv__create_schema_resp__pack(&resp, buf);
    drpc_resp->body.len = resp_len;
    drpc_resp->body.data = buf;
  }
  kv__create_schema_req__free_unpacked(req, &alloc.alloc);
}
void process_drpc_request(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *ctx)
{
  switch (drpc_req->method)
  {
  case DRPC_METHOD_CREATE_SCHEMA:
    create_schema(drpc_req, drpc_resp, ctx);
    break;
  default:
    break;
  }
}