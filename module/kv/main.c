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
#include "../drpc/drpc.pb-c.h"
#include "../mgmt/mgmt.pb-c.h"
#include "../mgmt/mgmt.h"
#include "../utils/log.h"

#define MAXEVENTS 64
#define DEFAULT_LOCAL_SOCKET_FILE "/tmp/drpc_socket.sock"
inline static void remove_drpc_socket_file(const char *name)
{
  if (access(name, F_OK) == 0)
  {
    remove(name);
  }
}
static int
make_socket_non_blocking(int sfd)
{
  int flags, s;

  flags = fcntl(sfd, F_GETFL, 0);
  if (flags == -1)
  {
    perror("fcntl");
    return -1;
  }

  flags |= O_NONBLOCK;
  s = fcntl(sfd, F_SETFL, flags);
  if (s == -1)
  {
    perror("fcntl");
    return -1;
  }

  return 0;
}

int main(int argc, char *argv[])
{
 
  log_init(NULL);
  int sfd, s;
  int efd;
  struct epoll_event event;
  struct epoll_event *events;

  efd = epoll_create1(0);
  if (efd == -1)
  {
    perror("epoll_create");
    abort();
  }
  struct mgmt *mgmt_ptr = mgmt_alloc(argv[1],argv[2]);

  remove_drpc_socket_file(DEFAULT_LOCAL_SOCKET_FILE);
  struct drpc *listener_ctx = drpc_listen(DEFAULT_LOCAL_SOCKET_FILE, mgmt_drpc_handlers[0].handler);
  sfd = listener_ctx->fd;
  event.data.fd = listener_ctx->fd;
  event.events = EPOLLIN | EPOLLET;
  s = epoll_ctl(efd, EPOLL_CTL_ADD, sfd, &event);
  if (s == -1)
  {
    perror("epoll_ctl");
    abort();
  }

  /* Buffer where events are returned */
  events = calloc(MAXEVENTS, sizeof event);
 
  /* The event loop */
  while (1)
  {
    int n, i;

    n = epoll_wait(efd, events, MAXEVENTS, -1);
    for (i = 0; i < n; i++)
    {
      if ((events[i].events & EPOLLERR) ||
          (events[i].events & EPOLLHUP) ||
          (!(events[i].events & EPOLLIN)))
      {
        close(events[i].data.fd);
        log_info("fd=%d closed", events[i].data.fd);
        continue;
      }

      else if (sfd == events[i].data.fd)
      {
        struct sockaddr in_addr;
        char hbuf[NI_MAXHOST], sbuf[NI_MAXSERV];
        socklen_t in_len = sizeof in_addr;
        struct drpc *session_ctx = drpc_accept(listener_ctx);
        if (session_ctx == NULL)
        {
          log_info("session_ctx=%p is nil", session_ctx);
          break;
        }
        int infd = session_ctx->fd;
        if (infd == -1)
        {
          if ((errno == EAGAIN) ||
              (errno == EWOULDBLOCK))
          {
            break;
          }
        }

        s = getnameinfo(&in_addr, in_len,
                        hbuf, sizeof hbuf,
                        sbuf, sizeof sbuf,
                        NI_NUMERICHOST | NI_NUMERICSERV);
        if (s == 0)
        {
          log_info("Accepted connection on descriptor %d (host=%s, port=%s)", infd, hbuf, sbuf);
        }

        event.data.fd = infd;
        event.data.ptr = session_ctx;
        event.events = EPOLLIN | EPOLLET;
        s = epoll_ctl(efd, EPOLL_CTL_ADD, infd, &event);
        if (s == -1)
        {
          epoll_ctl(efd, EPOLL_CTL_DEL, infd, NULL);
          log_info("fd=%d closed", infd);
        }

        continue;
      }
      else
      {
        struct drpc *session_ctx = events[i].data.ptr;
        Drpc__Call *incoming;
        int result = drpc_recv_call(session_ctx, &incoming);
        Drpc__Response *resp = drpc_response_create(incoming);
        log_info("enter handler=%p", session_ctx->handler);
        session_ctx->handler(incoming, resp,mgmt_ptr);
        drpc_send_response(session_ctx, resp);
        drpc_response_free(resp);
      }
    }
  }

  free(events);

  close(sfd);

  return EXIT_SUCCESS;
}
