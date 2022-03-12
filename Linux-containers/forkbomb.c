/* -*- compile-command: "gcc -Wall -Werror -static forkbomb.c -o forkbomb" -*- */
#include <stdio.h>
#include <unistd.h>
#include <errno.h>

int main(int argc, char **argv)
{
    switch (fork())
    {
    case -1:
        fprintf(stderr, "++ couldn't even fork once: %m\n");
        return 1;
    case 0:
        while (1)
        {
            switch (fork())
            {
            case -1:
                break;
            case 0:
                fprintf(stderr, "++ successful fork.\n");
                break;
            default:
                break;
            }
        }
        break;
    default:
        while (1)
            sleep(1);
        break;
    }
    return 0;
}