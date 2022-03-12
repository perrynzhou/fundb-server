/*************************************************************************
    > File Name: server.h
  > Author:perrynzhou 
  > Mail:perrynzhou@gmail.com 
  > Created Time: Wednesday, September 09, 2020 AM08:34:22
 ************************************************************************/

#ifndef _SERVER_H
#define _SERVER_H
#include "../drpc/drpc_util.h"
#include <ev.h>
const char *server_types[] = {
    "drpc",
    "sync",
};
typedef enum {
    DRPC_SERVER_TYPE=0,
    SYNC_SERVER_TYPE,
}server_type_t;

typedef struct  {
  ev_io io;
  struct drpc *listener;
  int  server_type;
  drpc_handler_func handler;
  char *socket;
  void *ctx;
}server_t;
//  server impl
server_t *server_alloc(int server_type,const char *socket_name,drpc_handler_func handler,void *ctx);

void   server_start(server_t *srv);
void server_free(server_t *srv);
#endif
