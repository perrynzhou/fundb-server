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
#include "log.h"
#include "kv_db.h"
#include "store.h"
#include "server.h"
#define MAXEVENTS 64


void *thread_cb(void *ctx) {
   server_t *srv = (server_t *)ctx;
   server_start(srv);
   return NULL;
}
int main(int argc, char *argv[])
{
 
  int n=1;
  kv_db_t *db = kv_db_alloc(argv[1],argv[2]);
  pthread_t *threads = calloc(n,sizeof(pthread_t));
  assert(threads !=NULL);
  log_init(NULL);
  kv_db_t *db = kv_db_alloc(argv[1],argv[2]);
  assert(db !=NULL);
  server_t *drpc_server = server_alloc(DRPC_SERVER_TYPE,1,kv_drpc_handlers[0].handler,db);
  pthread_create(&threads[0],NULL,&thread_cb,drpc_server);
  for(int i=0;i<n;i++) {
    pthread_join(threads[i],NULL);
  }
  free(threads !=NULL);
  server_free(drpc_server);
  return EXIT_SUCCESS;
}
