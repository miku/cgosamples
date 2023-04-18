libggml.so: ggml.c
	# cc -c -fPIC  -o ggml.o ggml.c
	# gcc -shared -o libggml.so ggml.o -lm
	gcc -shared -fPIC -Wall -g ggml.c -L. -lm -o libggml.so

main: main.go libggml.so
	go build -o main main.go

.PHONY: clean
clean:
	rm -f ggml.o
	rm -f libggml.so
	rm -f main
