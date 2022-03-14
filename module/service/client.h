/*************************************************************************
  > File Name: client.h
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Wednesday, September 09, 2020 AM08:34:22
 ************************************************************************/

#ifndef _CLIENT_H
#define _CLIENT_H
#include "../drpc/drpc.pb-c.h"
#include "store.h"
#include "kv_db.h"
#include <stdlib.h>
#include <ev.h>
typedef struct
{
    ev_io io;
    int cfd;
    kv_db_t *db_ctx;
    struct drpc *session_ctx;
      struct ev_loop *loop;
} client_t;

// client impl
client_t *client_alloc(int fd);
void client_free(client_t *c);

#endif
