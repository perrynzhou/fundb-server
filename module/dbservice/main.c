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
#include "dbservice.h"
#include "server.h"
#include "conf.h"
#include "schmea_meta.h"
#define MAXEVENTS 64
#define DEFAULT_PORT 8080
#define MAX_CONN 16
#define MAX_EVENTS 32
#define BUF_SIZE 16
#define MAX_LINE 256

static void *server_thread_cb(void *arg)
{
  server_t *srv = (server_t *)arg;
  log_info("run sock:%s", srv->socket);
  server_start(srv);
  return NULL;
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

  json_t *json_thread_num = conf_search(conf, "thread_num");
  json_t *json_db_name = conf_search(conf, "db_name");
  json_t *json_db_path = conf_search(conf, "db_path");

  char *db_name = json_string_value(json_db_name);
  char *db_path = json_string_value(json_db_path);
  int thread_num = json_integer_value(json_thread_num);
  log_init(NULL);
  kv_db_t *db = kv_db_alloc(db_name, db_path);
  kv_schema_t *schmea_meta = kv_schema_alloc(sys_schmea_meta_name, db, false);
  db->schmea_meta_cache = schema_cache_load(sys_schmea_meta_name, db);
  assert(db != NULL);
  int n = thread_num;
  pthread_t threads[thread_num];
  for (int i = 0; i < n; i++)
  {
    server_t *srv = server_alloc(DRPC_SERVER_TYPE, i, dbservice_drpc_handlers[0].handler, db);
    pthread_create(&threads[i], NULL, server_thread_cb, srv);
  }
  for (int i = 0; i < n; i++)
  {
    pthread_join(&threads[i], NULL);
  }

  return 0;
}
