/* -*- compile-command: "gcc -Wall -Werror -static try_regain_cap.c -o try_regain_cap" -*- */
#include <linux/capability.h>
#include <sys/prctl.h>
#include <stdio.h>

int main(int argc, char **argv)
{
    if (prctl(PR_CAPBSET_READ, CAP_MKNOD, 0, 0, 0))
    {
        fprintf(stderr, "++ have CAP_MKNOD\n");
    }
    else
    {
        fprintf(stderr, "++ don't have CAP_MKNOD\n");
    }
    return 0;
}
