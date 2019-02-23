.PHONY: all clean test


all: gomain/gomain goclib.a cmain/cmain

clean:
	rm -f lib/goclib/goclib.a	 gomain/gomain
	

gomain/gomain: gomain/gomain.go
	cd gomain && go build

goclib.a: lib/goclib/goclib.a
lib/goclib/goclib.a: lib/goclib/goclib.go
	cd lib/goclib && go build -buildmode=c-archive goclib.go

cmain/cmain: cmain/main.cpp lib/goclib/goclib.a lib/goclib/goclib.h
	cd cmain && gcc -pthread -o cmain -I ../lib/goclib/ main.cpp ../lib/goclib/goclib.a

test: cmain/cmain gomain/gomain
	time ./gomain/gomain
	@echo "-------------------"
	time ./cmain/cmain
