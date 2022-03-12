/* Local Variables: */
/* compile-command: "gcc -Wall -Werror -static enumerate_net_devs.c \*/
/*                   -o enumerate_net_devs" */
/* End: */
#include <stdio.h>
#include <net/if.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/ioctl.h>

int main(int argc, char **argv)
{
    int sock = socket(PF_LOCAL, SOCK_SEQPACKET, 0);
    for (size_t i = 0; i < 100; i++)
    {
        struct ifreq req = {.ifr_ifindex = i};
        if (!ioctl(sock, SIOCGIFNAME, &req))
            printf("%3lu: %s\n", i, req.ifr_name);
    }
    return 0;
}