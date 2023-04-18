main: main.go libggml.so
	go build -o main main.go

libggml.so: ggml.c
	# https://explainshell.com/explain?cmd=gcc+-shared+-fPIC+-Wall+-g+ggml.c+-L.+-lm+-o+libggml.so
	gcc -shared -fPIC -Wall -g ggml.c -L . -lm -o libggml.so


.PHONY: clean
clean:
	rm -f ggml.o
	rm -f libggml.so
	rm -f main
