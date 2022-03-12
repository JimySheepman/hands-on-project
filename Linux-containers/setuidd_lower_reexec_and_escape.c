/* -*- compile-command: "gcc -Wall -Werror -static setuidd_lower_reexec_and_escape.c -o setuidd_lower_reexec_and_escape" -*- */
#define _GNU_SOURCE
#include <stdio.h>
#include <unistd.h>
#include <sys/ioctl.h>

int main(int argc, char **argv)
{
    if (argc == 1)
    {
        if (setresuid(99, 99, 99))
        {
            fprintf(stderr, "++ setresuid failed: %m\n");
            return 1;
        }
        if (execve(argv[0], (char *[]){argv[0], "-", 0}, NULL))
        {
            fprintf(stderr, "++ execve failed: %m\n");
            return 1;
        }
    }
    else
    {
        uid_t a, b, c = 0;
        getresuid(&a, &b, &c);
        fprintf(stderr, "++ we're %u/%u/%u.\n", a, b, c);
        if (ioctl(STDIN_FILENO, TIOCSTI, "!"))
        {
            fprintf(stderr, "++ ioctl failed: %m\n");
            return 1;
        }
    }
}
