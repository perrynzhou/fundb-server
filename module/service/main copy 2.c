/*************************************************************************
  > File Name: server.c
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: 六 11/20 13:36:21 2021
 ************************************************************************/

#include <stdio.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/epoll.h>
#include <errno.h>
#include <pthread.h>
#include <assert.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/epoll.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/epoll.h>
#include <errno.h>

#include "log.h"
#include "kv_db.h"
#include "store.h"
#include "server.h"
#include "conf.h"
#define MAXEVENTS 64
#define DEFAULT_PORT 8080
#define MAX_CONN 16
#define MAX_EVENTS 32
#define BUF_SIZE 16
#define MAX_LINE 256
/*
 * register events of fd to epfd
 */
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

void server_run(const char *socket, kv_db_t *db)
{
  int i;
  int n;
  int epfd;
  int nfds;
  int listen_sock;
  int conn_sock;
  int socklen;

  char buf[BUF_SIZE];
  struct epoll_event events[MAX_EVENTS];

  struct epoll_event event;
  struct drpc *listener_ctx = drpc_listen(socket, kv_drpc_handlers[0].handler);
  listen_sock = listener_ctx->fd;

  setnonblocking(listen_sock);
  epfd = epoll_create(1);
  epoll_ctl_add(epfd, listen_sock, EPOLLIN | EPOLLOUT | EPOLLET);

  for (;;)
  {
    nfds = epoll_wait(epfd, events, MAX_EVENTS, -1);
    for (i = 0; i < nfds; i++)
    {

      if ((events[i].events & EPOLLERR) ||
          (events[i].events & EPOLLHUP) ||
          (!(events[i].events & EPOLLIN)))
      {
        close(events[i].data.fd);
        log_info("fd=%d closed", events[i].data.fd);
        continue;
      }
      else if (listen_sock == events[i].data.fd)
      {
        struct sockaddr in_addr;
        char hbuf[1024], sbuf[1024];
        socklen_t in_len = sizeof in_addr;
        struct drpc *session_ctx = drpc_accept(listener_ctx);
        if (session_ctx == NULL || session_ctx->fd == -1)
        {
          log_info("session_ctx=%p is nil", session_ctx);
          break;
        }
        event.data.fd = session_ctx->fd;
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
        log_info("enter handler=%p", session_ctx->handler);
        session_ctx->handler(incoming, resp, db);
        drpc_send_response(session_ctx, resp);
      }
    }
  }
}

int main(int argc, char *argv[])
{

  char *conf_file = argv[1];
  if (conf_file == NULL)
  {
    return -1;
  }
  conf_t *conf = conf_alloc(conf_file);
  assert(conf != NULL);

  json_t *json_db_name = conf_search(conf, "db_name");
  json_t *json_db_path = conf_search(conf, "db_path");

  char *db_name = json_string_value(json_db_name);
  char *db_path = json_string_value(json_db_path);

  log_init(NULL);
  kv_db_t *db = kv_db_alloc(db_name, db_path);
  assert(db != NULL);
  server_run("/tmp/drpc.sock", db);
  /*
    server_t *srv = server_alloc(DRPC_SERVER_TYPE, kv_drpc_handlers[0].handler, db);
    server_start(srv);

    server_free(srv);
    */

  return 0;
}
