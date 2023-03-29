fmt:
	clang-format --style WebKit -i main.c

main.s: main.c
	cc -S main.c

main.o: main.c
	# Compile or assemble the source files, but do not link.  The linking stage
	# simply is not done.  The ultimate output is in the form of an object file
	# for each source file.
	cc -c main.c

main: main.o
	# ld -o main main.o /lib/x86_64-linux-gnu/crt1.o -lc --build-id -m elf_x86_64 -e _start -dynamic-linker /lib64/ld-linux-x86-64.so.2
	ld -o main main.o /lib/x86_64-linux-gnu/crt1.o -lc -dynamic-linker /lib64/ld-linux-x86-64.so.2

main.stripped.o: main.o
	strip -o main.stripped.o main.o

compile.log:
	strace -o compile.log -s 1024 -f -e trace=/exec gcc -o main main.c

.PHONY: clean
clean:
	rm -f main
	rm -f main.s
	rm -f main.o
	rm -f main.stripped.o
