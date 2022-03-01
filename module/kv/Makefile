all:
	rm -rf test*
	gcc -g -std=gnu11  -w -O0 hashfn.h hashfn.c dict.h dict.c kv_db.h kv_db.c  kv_db_test.c -o test_db -lwiredtiger
