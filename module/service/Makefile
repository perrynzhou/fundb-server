all:
	rm -rf test*
	gcc -g -std=gnu11  -w -O0 hashfn.h hashfn.c dict.h dict.c kv_db.h kv_db.c  kv_db_test.c -o test_db -lwiredtiger
	gcc -g -std=gnu11  -w -O0 ../drpc/drpc.pb-c.h ../drpc/drpc.pb-c.c ../drpc/drpc_util.h ../drpc/drpc.h ../drpc/drpc.c \
	 hashfn.h hashfn.c dict.h dict.c kv_db.h kv_db.c  kv.pb-c.h kv.pb-c.c \
	log.h log.c util.h util.c  store.h store.c  client.h client.c conf.h conf.c server.h server.c main.c -o test_server -lwiredtiger -lprotobuf-c -lev -lpthread -ljansson

