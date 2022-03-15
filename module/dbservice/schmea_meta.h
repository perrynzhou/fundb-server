/*************************************************************************
    > File Name: schmea_cache.h
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Tue 15 Mar 2022 07:46:54 AM UTC
 ************************************************************************/

#ifndef _SCHMEA_META_CACHE_H
#define _SCHMEA_META_CACHE_H
#include <stdio.h>
#include "kv_db.h"
#include "dict.h"
typedef struct
{
  uint32_t key_count;
  bool is_active;
} schema_meta_t;
static const char *schmea_meta_name = "schema_meta";
int schmea_meta_assign(schema_meta_t *meta,uint32_t key_count,bool is_active);
dict_t *schema_cache_load(const char *schema_name, kv_db_t *db);
int schmea_cache_add(dict_t *cache, kv_db_t *db, char *schema_name, const char *key, void *val, size_t val_sz);
int schmea_cache_del(dict_t *cache, kv_db_t *db, char *schema_name, const char *key);

#endif
