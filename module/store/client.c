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
#include <stdlib.h>
static void client_cb(EV_P_ ev_io *w, int revents)
{
    client_t *c = (struct client_t *)w;
    struct drpc *session_ctx = (struct drpc *)client->session_ctx;
    Drpc__Request *incoming;
    int result = drpc_recv_call(session_ctx, &incoming);
    if (result != -1)
    {
        kv_db_t *db_ctx = c->db_ctx;
        Drpc__Response *resp = drpc_response_create(incoming);
        log_info("enter handler=%p", session_ctx->handler);
        session_ctx->handler(incoming, resp, db_ctx);
        drpc_send_response(session_ctx, resp);
        drpc_response_free(resp);
        return;
    }
    ev_io_stop(EV_A_ & c->io);
    close(c->fd);
    client_free(c);
}
client_t *client_alloc(int fd)
{
    client_t *c = realloc(NULL, sizeof(struct client_t));
    c->fd = fd;
    // client->server = server;
    setnonblock(c->fd);
    ev_io_init(&c->io, client_cb, c->fd, EV_READ);
    return client;
}
void client_free(client_t *c)
{
    // todo
}