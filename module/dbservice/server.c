/*************************************************************************
    > File Name: server.C
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Wednesday, September 09, 2020 AM08:34:22
 ************************************************************************/

#include "server.h"
#include "../drpc/drpc.h"
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
#include <ev.h>
static void epoll_ctl_add(int epfd, int fd, uint32_t events)
{
  struct epoll_event ev;
  ev.events = events;
  ev.data.fd = fd;
  if (epoll_ctl(epfd, EPOLL_CTL_ADD, fd, &ev) == -1)
  {
    perror("epoll_ctl()\n");
    exit(1);
  }
}

static int setnonblocking(int sockfd)
{
  if (fcntl(sockfd, F_SETFD, fcntl(sockfd, F_GETFD, 0) | O_NONBLOCK) ==
      -1)
  {
    return -1;
  }
  return 0;
}

inline static void remove_socket(const char *name)
{
  if (access(name, F_OK) == 0)
  {
    remove(name);
  }
}
server_t *server_alloc(int server_type, int id, drpc_handler_func handler, void *ctx)
{
  server_t *srv = calloc(1, sizeof(server_t));
  assert(srv != NULL);
  char buffer[256] = {'\0'};
  snprintf(&buffer, 256, "/tmp/%s_%d.sock", server_type_names[server_type], id);
  srv->socket = strdup(&buffer);
  remove_socket(srv->socket);
  struct drpc *listener = drpc_listen(srv->socket, handler);
  srv->sfd = listener->fd;
  log_info("active fd=%d,socket=%s", srv->sfd, srv->socket);
  srv->server_type = server_type;
  srv->listener = listener;
  srv->db_ctx = (kv_db_t *)ctx;
  return srv;
}

void server_start(server_t *srv)
{
  int epfd;
  int nfds;
  int listen_sock;
  int max_event = 4096;
  struct epoll_event events[max_event];

  struct epoll_event event;
  listen_sock = srv->listener->fd;

  kv_db_t *db = (kv_db_t *)srv->db_ctx;
  setnonblocking(listen_sock);
  epfd = epoll_create(1);
  epoll_ctl_add(epfd, listen_sock, EPOLLIN | EPOLLOUT | EPOLLET);

  for (;;)
  {
    nfds = epoll_wait(epfd, events, max_event, -1);
    for (int i = 0; i < nfds; i++)
    {

      if ((events[i].events & EPOLLERR) ||
          (events[i].events & EPOLLHUP) ||
          (!(events[i].events & EPOLLIN)))
      {
        close(events[i].data.fd);
        continue;
      }
      else if (listen_sock == events[i].data.fd)
      {
        struct drpc *session_ctx = drpc_accept(srv->listener);
        if (session_ctx == NULL || session_ctx->fd == -1)
        {
          log_info("session_ctx=%p is nil", session_ctx);
          break;
        }

        event.data.ptr = session_ctx;
        event.events = EPOLLIN | EPOLLET;
        int s = epoll_ctl(epfd, EPOLL_CTL_ADD, session_ctx->fd, &event);
        if (s == -1)
        {
          epoll_ctl(epfd, EPOLL_CTL_DEL, session_ctx->fd, NULL);
          log_info("fd=%d closed", session_ctx->fd);
        }

        continue;
      }

      else
      {
        struct drpc *session_ctx = events[i].data.ptr;
        Drpc__Request *incoming;
        int result = drpc_recv_call(session_ctx, &incoming);
        Drpc__Response *resp = drpc_response_create(incoming);
        log_info("enter handler=%p,fd=%d", session_ctx->handler, session_ctx->fd);
        session_ctx->handler(incoming, resp, db);
        drpc_send_response(session_ctx, resp);
        drpc_response_free(resp);
      }
    }
  }
}

void server_free(server_t *srv)
{
  if (srv != NULL)
  {
    if (srv->sfd != NULL)
    {
      close(srv->sfd);
    }
    if (srv->socket != NULL)
    {
      free(srv->socket);
    }
    free(srv);
    srv = NULL;
  }
}
