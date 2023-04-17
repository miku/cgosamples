.PHONY: fmt
fmt:
	clang-format -style WebKit -i *.cpp
	clang-format -style WebKit -i *.h

librect.o: rect.cpp
	g++ -c -o librect.o rect.cpp

librect.so: librect.o
	gcc -shared -o librect.so librect.o

.PHONY: clean
clean:
	rm librect.o
	rm librect.so
