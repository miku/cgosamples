.PHONY: all
all: fmt lcg_rand randbits main

lcg_rand: lcg_rand.c
	cc -O3 -o lcg_rand lcg_rand.c

.PHONY: fmt
fmt:
	clang-format --style Chromium -i *.c

randbits: randbits.go
	go build -o randbits randbits.go

main: main.go
	# CGO_CPPFLAGS="-O0 -fno-inline" go build -gcflags '-N -l' -o main main.go
	go build -o main main.go

.PHONY: clean
clean:
	rm -f main
	rm -f lcg_rand
	rm -f randbits

