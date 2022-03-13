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
#include "log.h"
#include "kv_db.h"
#include "store.h"
#include "server.h"
#include "conf.h"
#define MAXEVENTS 64

typedef struct
{
  server_t *srv;
  pthread_t thread;
} server_thread_t;

void *server_thread_cb(void *ctx)
{
  server_thread_t *srv_thread = (server_thread_t *)ctx;
  server_t *srv = srv_thread->srv;
  server_start(srv_thread->srv);
  return NULL;
}
server_thread_t *server_thread_alloc(server_t *srv)
{
  server_thread_t *srv_thread = calloc(1, sizeof(server_thread_t));
  assert(srv_thread != NULL);

  srv_thread->srv = srv;
  pthread_create(&srv_thread->thread, NULL, &server_thread_cb, srv_thread);

  return srv_thread;
}
void server_thread_free(server_thread_t *srv_thread)
{
  if (srv_thread != NULL)
  {
    pthread_join(srv_thread->thread, NULL);
    free(srv_thread);
    srv_thread = NULL;
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
  json_t *json_thread_num = conf_search(conf, "thread_num");

  int  thread_num= json_integer_value(json_thread_num);
  char *db_name = json_string_value(json_db_name);
  char *db_path = json_string_value(json_db_path);
  
  log_init(NULL);
  kv_db_t *db = kv_db_alloc(db_name, db_path);
  assert(db != NULL);
  server_thread_t **srv_threads = calloc(thread_num, sizeof(server_thread_t *));

  for (int i = 0; i < thread_num; i++)
  {
    server_t *drpc_server = server_alloc(DRPC_SERVER_TYPE, i, kv_drpc_handlers[0].handler, db);
    srv_threads[i] = server_thread_alloc(drpc_server);
  }
  for (int i = 0; i < thread_num; i++)
  {
    server_free(srv_threads[i]->srv);
    server_thread_free(srv_threads[i]);
  }
  free(srv_threads);
  return 0;
}
