/*************************************************************************
  > File Name: drpc_util.h
  > Author:perrynzhou 
  > Mail:perrynzhou@gmail.com 
  > Created Time: Thu 30 Dec 2021 05:40:18 PM CST
 ************************************************************************/

#ifndef _DRPC_UTIL_H
#define _DRPC_UTIL_H
#include "./drpc.h"
#define PROTO_ALLOCATOR_INIT(self)  \
  {                                 \
    .alloc.alloc = daos_drpc_alloc, \
    .alloc.free = daos_drpc_free,   \
    .alloc.allocator_data = &self   \
  }
struct drpc_handler
{
  int module_id;
  drpc_handler_func handler;
};
#endif
