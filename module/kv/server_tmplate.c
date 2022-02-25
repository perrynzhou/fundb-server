/*************************************************************************
  > File Name: server.c
  > Author:perrynzhou 
  > Mail:perrynzhou@gmail.com 
  > Created Time: 六 11/20 13:36:21 2021
 ************************************************************************/

#include <stdio.h>
#include "drpc.pb-c.h"
#include "mgmt.pb-c.h"
#include "mgmt.h"
#include "drpc.h"
int main(int argc, char *argv[])
{

  struct drpc *listener_ctx = drpc_listen("/tmp/drpc_socket.sock", mgmt_drpc_handlers[0].handler);
  while (1)
  {
    struct drpc *session_ctx = drpc_accept(listener_ctx);
    Drpc__Call *incoming;
    int result = drpc_recv_call(session_ctx, &incoming);
    if (result != 0)
    {
      Drpc__Response *resp = drpc_response_create(incoming);
      session_ctx->handler(incoming, resp);
    }
  }
}
