/*************************************************************************
    > File Name: dbservice_help.c
  > Author:perrynzhou 
  > Mail:perrynzhou@gmail.com 
  > Created Time: Thu 17 Mar 2022 01:54:59 AM UTC
 ************************************************************************/

#include <stdio.h>
#include "dbservice_help.h"
// dbservice__create_schema_req__free_unpacked

// dbservice__create_schema_resp__pack

void dbservice_finish_response(Drpc__Request *drpc_req, Drpc__Response *drpc_resp, void *request, void *response, void *alloc)
{
  size_t resp_len = dbservice_response_get_packed_size(drpc_req, response);
  uint8_t *buf;
  D_ALLOC(buf, resp_len);
  if (buf != NULL)
  {
    dbservice_response_pack(drpc_req, response, buf);
    drpc_resp->body.len = resp_len;
    drpc_resp->body.data = buf;
  }
  dbservice_request_free_unpacked(drpc_req, request, alloc);
}