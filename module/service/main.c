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

  server_t *srv = server_alloc(DRPC_SERVER_TYPE, kv_drpc_handlers[0].handler, db);
  server_start(srv);

  server_free(srv);

  return 0;
}
