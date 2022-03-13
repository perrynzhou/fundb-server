/*************************************************************************
    > File Name: server.C
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Wednesday, September 09, 2020 AM08:34:22
 ************************************************************************/

#include "server.h"
#include "client.h"
#include "../drpc/drpc_util.h"
#include <stdlib.h>
#include <string.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/epoll.h>
#include <errno.h>

static void not_blocked(EV_P_ ev_periodic *w, int revents)
{
    fprintf(stdout, "i'm not blocked\n");
}
static void server_cb(EV_P_ ev_io *w, int revents)
{

    server_t *srv = (server_t *)w;
    struct drpc *session_ctx = NULL;
    while ((session_ctx = drpc_accept(srv->listener))!=NULL)
    {
        if (session_ctx->fd !=-1)
        {
            log_info("*****request on %s******",srv->socket);
            log_info("server sfd=%d,accepted a client=%d",srv->listener->fd,session_ctx->fd);
            client_t *c = client_alloc(session_ctx->fd);
            c->session_ctx = session_ctx;
            c->db_ctx = srv->db_ctx;
            ev_io_start(EV_A_ & c->io);
        }
    }
}

inline static void remove_socket(const char *name)
{
    if (access(name, F_OK) == 0)
    {
        remove(name);
    }
}
server_t *server_alloc(int server_type, int id,drpc_handler_func handler, void *ctx)
{
    server_t *srv = calloc(1, sizeof(server_t));
    assert(srv != NULL);
    char buffer[256] = {'\0'};
    snprintf(&buffer, 256, "/tmp/%s_%d.sock", server_type_names[server_type],id);
    srv->socket = strdup(&buffer);
    remove_socket(srv->socket);
    struct drpc *listener = drpc_listen(srv->socket, handler);
    srv->fd = listener->fd;
    log_info("active fd=%d",srv->fd);
    srv->server_type = server_type;
    srv->listener = listener;
    srv->db_ctx = (kv_db_t *)ctx;
    return srv;
}

void server_start(server_t *srv)
{
    EV_P = ev_loop_new(EVFLAG_AUTO);
    ev_io_init(&srv->io, server_cb, srv->fd, EV_READ);
    ev_io_start(EV_A_ & srv->io);
    log_info("start socket at ::%s  on %ld", srv->socket, pthread_self());
    ev_loop(EV_A_ 0);

    close(srv->fd);
    return EXIT_SUCCESS;
}

void server_free(server_t *srv)
{
    if (srv != NULL)
    {
        if (srv->fd != NULL)
        {
            close(srv->fd);
        }
        if (srv->socket != NULL)
        {
            free(srv->socket);
        }
        free(srv);
        srv = NULL;
    }
}
