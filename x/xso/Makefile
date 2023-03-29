fmt:
	clang-format --style WebKit -i main.c

main.o:
	# Compile or assemble the source files, but do not link.  The linking stage
	# simply is not done.  The ultimate output is in the form of an object file
	# for each source file.
	cc -c main.c
