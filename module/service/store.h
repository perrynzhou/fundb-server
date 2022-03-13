/*************************************************************************
  > File Name: store.h
  > Author:perrynzhou 
  > Mail:perrynzhou@gmail.com 
  > Created Time: 二  3/ 8 11:10:56 2022
 ************************************************************************/

#ifndef _STORE_H
#define _STORE_H
#include <stdio.h>
#include <stdbool.h>
#include "../drpc/drpc.pb-c.h"
#include "../drpc/drpc_util.h"
#include "log.h"
#include "kv.pb-c.h"
void process_drpc_request(Drpc__Request *drpc_req, Drpc__Response *drpc_resp,void *ctx);

 static struct drpc_handler kv_drpc_handlers[] = {
    {.module_id = 0,
     .handler = process_drpc_request}};
#endif
