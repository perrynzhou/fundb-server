/*************************************************************************
  > File Name: client.c
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Wednesday, September 09, 2020 AM08:34:22
 ************************************************************************/
#include "client.h"
#include "../drpc/drpc.pb-c.h"
#include "store.h"
#include "kv_db.h"
#include "util.h"
#include "log.h"
#include <stdlib.h>
static void client_cb(  struct ev_loop *loop, ev_io *w, int revents)
{
    client_t *c = (struct client_t *)w;
    struct drpc *session_ctx = (struct drpc *)c->session_ctx;
    Drpc__Request *incoming;
    int result = drpc_recv_call(session_ctx, &incoming);
    log_info("client cb result=%d", result);
    if (result != -1)
    {
        kv_db_t *db_ctx = c->db_ctx;
        Drpc__Response *resp = drpc_response_create(incoming);
        log_info("enter handler=%p", session_ctx->handler);
        session_ctx->handler(incoming, resp, db_ctx);
        drpc_send_response(session_ctx, resp);
        drpc_response_free(resp);
        drpc_call_free(incoming);
        ev_io_stop(c->loop,& c->io);
        client_free(c);
    }
}
client_t *client_alloc(int cfd)
{
    client_t *c = calloc(1, sizeof(client_t));
    c->cfd = cfd;
    // client->server = server;
    setnonblock(c->cfd);
    ev_io_init(&c->io, client_cb, c->cfd, EV_READ);
    return c;
}
void client_free(client_t *c)
{
    if (c != NULL)
    {
        if (c->cfd != -1)
        {
            close(c->cfd);
        }
        free(c);
        c = NULL;
    }
    // todo
}