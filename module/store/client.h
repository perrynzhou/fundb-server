/*************************************************************************
  > File Name: client.h
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Wednesday, September 09, 2020 AM08:34:22
 ************************************************************************/

#ifndef _CLIENT_H
#define _CLIENT_H
#include <ev.h>
struct
{
    ev_io io;
    int fd;
    server_t *server;
    void *ctx;
} client_t;

// client impl
client_t *client_alloc(int fd);
void client_free(client_t *c);

#endif
