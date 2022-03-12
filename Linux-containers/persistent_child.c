/* -*- compile-command: "gcc -Wall -Werror -static persistent_child.c -o persistent_child" -*- */
#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

int main(int argc, char **argv)
{
    switch (fork())
    {
    case -1:
        fprintf(stderr, "++ fork failed: %m\n");
        return 1;
    case 0:;
        int fd = 0;
        if ((fd = open("persistent_child.log",
                       O_CREAT | O_APPEND | O_WRONLY,
                       S_IRUSR | S_IWUSR)) == -1)
        {
            fprintf(stderr, "++ open failed: %m\n");
            return 1;
        }
        size_t count = 0;
        while (count < 100)
        {
            if (dprintf(fd, "%lu\n", count++) < 0)
            {
                fprintf(stderr, "++ dprintf failed: %m\n");
                close(fd);
                return 1;
            }
            sleep(1);
        }
        close(fd);
        return 0;
    default:
        sleep(2);
        return 0;
    }
}
