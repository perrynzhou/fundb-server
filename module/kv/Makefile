all:
	rm -rf test_db
	gcc -g -std=gnu11  hashfn.h hashfn.c dict.h dict.c kv_db.h kv_db.c  kv_db_test.c -o test_db -lwiredtiger