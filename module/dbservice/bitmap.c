/*************************************************************************
    > File Name: bitmap.c
  > Author:perrynzhou
  > Mail:perrynzhou@gmail.com
  > Created Time: Fri 01 Apr 2022 10:56:39 AM UTC
 ************************************************************************/

#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include "bitmap.h"

inline int bitmap_getbit(bitmap_t *bmap, uint64_t idx)
{
  if (bmap != NULL)
  {
    return (bmap->bits[idx >> 3] >> (7 - (idx % 8))) & 0x1;
  }
  return -1;
}

inline void bitmap_setbit(bitmap_t *bmap, uint64_t idx)
{
  if (bmap != NULL)
  {
    unsigned char byte = bmap->bits[idx >> 3];
    unsigned char byte_off = 7 - idx % 8;
    byte |= 1 << byte_off;
    bmap->bits[idx >> 3] = byte;
  }
}
int bitmap_init(bitmap_t *bmap, uint64_t cap)
{
  if (bmap != NULL)
  {
    bmap->bits = (unsigned char *)calloc(cap, sizeof(char));
    assert(bmap->bits != NULL);
    bmap->size = cap;
    return 0;
  }
  return -1;
}

bitmap_t *bitmap_alloc(uint64_t cap)
{
  bitmap_t *bmap = (bitmap_t *)calloc(1, sizeof(bitmap_t));
  assert(bmap != NULL);
  if (bitmap_init(bmap, cap) != 0)
  {
    free(bmap);
    bmap = NULL;
  }
  return bmap;
}

void bitmap_deinit(bitmap_t *bmap)
{
  if (bmap != NULL)
  {
    free(bmap->bits);
    bmap->bits = NULL;
  }
}
void bitmap_destroy(bitmap_t *bmap)
{
  bitmap_deinit(bmap);
  if (bmap != NULL)
  {
    free(bmap);
    bmap = NULL;
  }
}

int bitmap_dump(bitmap_t *bmap, const char *filename)
{
  int ret = -1;
  int fd = open(filename, O_RDWR | O_TRUNC);
  if (fd != -1 && bmap != NULL)
  {
    if (write(fd, bmap->bits, bmap->size) > 0)
    {
      ret = 0;
    }
  }
  if (fd != -1)
  {
    close(fd);
  }
  return ret;
}
int bitmap_load(bitmap_t *bmap, const char *filename)
{
  int ret = -1;
  if (bmap != NULL)
  {
    int fd = open(filename, O_RDWR | O_TRUNC);
    if (fd != -1)
    {
      struct stat st;
      stat(filename, &st);
      bitmap_init(bmap, st.st_size);
      read(fd, bmap->bits, st.st_size);
      close(fd);
      ret = 0;
    }
  }
  return ret;
}
