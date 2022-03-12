/* -*- compile-command: "gcc -Wall -Werror harmless_setuid.c -o harmless_setuid" -*- */
#define _GNU_SOURCE
#include <unistd.h>
#include <stdio.h>

int main (int argc, char **argv)
{
	uid_t a, b, c = 0;
	getresuid(&a, &b, &c);
	printf("I'm #%d/%d/%d\n", a, b, c);
	return 0;
}
