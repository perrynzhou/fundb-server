/*************************************************************************
    > File Name: schmea_meta_cache.c
  > Author:perrynzhou 
  > Mail:perrynzhou@gmail.com 
  > Created Time: Tue 15 Mar 2022 07:44:56 AM UTC
 ************************************************************************/

#include <stdio.h>
#include "kv_db.h"
#include "dict.h"
#include "log.h"
#include "schmea_meta.h"
#define SCHEMA_MAX_SIZE (65535)
int schmea_meta_assign(schema_meta_t *meta,uint32_t key_count,bool is_active)
{
  memset(meta,0,sizeof(*meta));
  meta->is_active=is_active;
  meta->key_count=key_count;
  return 0;
}
static int schema_cache_load_cb(void *ctx,void *key,void *val)
{
  dict_t *cache = (dict_t *)ctx;
  
  schema_meta_t *meta_ptr=(schema_meta_t *)dict_put(cache,(char *)key,val);
  log_info("cache::load key=%s,count=%d,active=%d",(char *)key,meta_ptr->key_count,meta_ptr->is_active);
  return (meta_ptr == NULL)?-1:0;
}

int schmea_cache_add(dict_t *cache,kv_db_t *db,char *schema_name,const char *key,void *val,size_t val_sz)
{
  schema_meta_t *meta_ptr = (schema_meta_t *)val;
  log_info("cache:add key=%s,count=%d,active=%d",(char *)key,meta_ptr->key_count,meta_ptr->is_active);

  int ret =kv_db_set(db,schema_name,key,strlen(key),val,val_sz);
  if(ret !=0) {
    return -1;
  }
  log_info("cache:add key=%s,count=%d,active=%d",(char *)key,meta_ptr->key_count,meta_ptr->is_active);
  void *ptr =dict_put(cache,(char *)key,val);
  return (ptr == NULL)?-1:0;
}
int schmea_cache_del(dict_t *cache,kv_db_t *db,char *schema_name,const char *key) {

  log_info("cache:del key=%s",(char *)key);
  int ret =kv_db_del(db,schema_name,key,strlen(key));
  if(ret !=0) {
    return ret;
  }
  log_info("cache:del key=%s",(char *)key);
  dict_del(cache,key,NULL);
  return 0;
}
dict_t *schema_cache_load(const char *meta_name,kv_db_t *db)
{
   kv_schema_t *meta = kv_db_fetch_schema(db,meta_name);
   dict_t *map = dict_create(SCHEMA_MAX_SIZE,NULL);
   kv_schmea_search(meta,schema_cache_load_cb,map,NULL,0);
   return map;

}

