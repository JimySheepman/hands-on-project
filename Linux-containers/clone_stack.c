/* -*- compile-command: "gcc -Wall -Werror clone_stack.c -o clone_stack" -*- */
#define _GNU_SOURCE
#include <sched.h>
#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define STACK_SIZE (1024 * 1024)

int child(void *_)
{
    int stack_value = 0;
    fprintf(stderr, "pre-execve, stack is ~%p\n", &stack_value);
    execve("./show_stack", (char *[]){",/show_stack", 0}, NULL);
    return 0;
}

int main(int argc, char **argv)
{
    void *stack = malloc(STACK_SIZE);
    clone(child, stack + STACK_SIZE, SIGCHLD, NULL);
    wait(NULL);
    return 0;
}
