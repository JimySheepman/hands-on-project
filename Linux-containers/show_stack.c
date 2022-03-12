/* -*- compile-command: "gcc -Wall -Werror -static show_stack.c -o show_stack" -*- */
#include <stdio.h>

int main(int argc, char **argv)
{
    int stack_value = 0;
    fprintf(stderr, "post-execve, stack is ~%p\n", &stack_value);
    return 0;
}
