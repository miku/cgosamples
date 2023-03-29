#include <stdio.h>
#include <stdlib.h>

void print_name(const char* name)
{
    printf("my name is %s\n", name);
}

int main(int argc, char* argv[])
{
    print_name("gopher");
    return 0;
}
