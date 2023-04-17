main: main.go librect.so
	go build -o main main.go

.PHONY: fmt
fmt:
	clang-format -style WebKit -i *.c
	clang-format -style WebKit -i *.cpp
	clang-format -style WebKit -i *.h
	clang-format -style WebKit -i *.hpp

rect.o: rect.cpp
	g++ -c -o rect.o rect.cpp

crect.o: rect.c
	g++ -c -o crect.o rect.c

librect.so: rect.o crect.o
	gcc -shared -o librect.so rect.o crect.o

.PHONY: clean
clean:
	rm -f rect.o
	rm -f crect.o
	rm -f librect.so
	rm -f main
