/*************************************************************************
  > File Name: store.c
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: 二  3/ 8 11:11:03 2022
 ************************************************************************/

#include <stdio.h>
#include "../drpc/drpc.pb-c.h"
#include "../drpc/drpc_util.h"
#include "log.h"
#include "kv.pb-c.h"
#include "store.h"
#include "kv_db.h"
enum drpc_kv_method
{
  DRPC_METHOD_CREATE_SCHEMA = 201,
};

static void create_schema(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *ctx)
{
  kv_db_t *db = (kv_db_t *)ctx;
  struct drpc_alloc alloc = PROTO_ALLOCATOR_INIT(alloc);
  void *ptr = kv__create_schema_req__unpack(NULL, drpc_req->body.len, drpc_req->body.data);
  Kv__CreateSchemaReq *req = kv__create_schema_req__unpack(&alloc.alloc, drpc_req->body.len, drpc_req->body.data);
  kv_schema_t *schema = kv_db_fetch_schema(db, req->name);
  bool done = false;
  if (NULL == schema)
  {
    schema = kv_schema_alloc(req->name, db, false);
    done = true;
  }
  Kv__CreateSchemaResp *resp = NULL;
  D_ALLOC_PTR(resp);
  kv__create_schema_resp__init(resp);
  resp->msg = req->name;
  if (schema != NULL && done)
  {
    resp->code = 0;
    resp->msg = "succ";
  }
  else if (schema != NULL && !done)
  {
    resp->code = -1;
    resp->msg = "exists";
  }
  else
  {
    resp->code = -1;
    resp->msg = "fail";
  }
  size_t resp_len = kv__create_schema_resp__get_packed_size(resp);
  uint8_t *buf;
  D_ALLOC(buf, resp_len);
  if (buf != NULL)
  {
    kv__create_schema_resp__pack(resp, buf);
    drpc_resp->body.len = resp_len;
    drpc_resp->body.data = buf;
  }
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