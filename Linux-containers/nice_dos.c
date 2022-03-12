/* -*- compile-command: "gcc -Wall -Werror -static nice_dos.c -o nice_dos" -*- */
#include <unistd.h>
#include <stdio.h>

int main (int argc, char **argv)
{
	if (nice(-10) == -1) {
		fprintf(stderr, "++ nice failed: %m\n");
		return 1;
	}
	if (execve("./busy_loop", (char *[]) { "./busy_loop", 0 }, NULL)) {
		fprintf(stderr, "++ execve failed: %m\n");
		return 1;
	}
}
