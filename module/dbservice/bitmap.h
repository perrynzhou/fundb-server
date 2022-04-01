/*************************************************************************
    > File Name: bitmap.h
  > Author:perrynzhou 
  > Mail:perrynzhou@gmail.com 
  > Created Time: Fri 01 Apr 2022 10:56:32 AM UTC
 ************************************************************************/

#ifndef _BITMAP_H
#define _BITMAP_H
#include <stdint.h>
typedef struct {         
    uint64_t size;       
    unsigned char* bits; 
} bitmap_t;

bitmap_t *bitmap_alloc(uint64_t cap);
int  bitmap_init(bitmap_t *bmap,uint64_t cap);
int bitmap_dump(bitmap_t *bmap,const char *filename);
int bitmap_load(bitmap_t *bmap,const char *filename);
int bitmap_getbit(bitmap_t *bmap, uint64_t idx);
void bitmap_setbit(bitmap_t *bmap, uint64_t idx);
void bitmap_deinit(bitmap_t *bmap);
void bitmap_destroy(bitmap_t *bmap);
#endif
