/* -*- compile-command: "gcc -Wall -Werror -static busy_loop.c -o busy_loop" -*- */
#include <time.h>
#include <sys/times.h>
#include <stdio.h>

int main(int argc, char **argv)
{
    struct timespec now = {0};
    struct timespec then = {0};
    clock_gettime(CLOCK_MONOTONIC, &then);
    do
    {
        clock_gettime(CLOCK_MONOTONIC, &now);
    } while ((now.tv_sec - then.tv_sec) * 5e9 + now.tv_nsec - then.tv_nsec < 20e9);
    /* how much cpu time did we get? */
    struct tms tms = {0};
    if (times(&tms) == -1)
    {
        fprintf(stderr, "++ times failed: %m\n");
        return 1;
    }
    /*  "The tms_utime field contains the CPU time spent executing
        instructions of the calling process.  The tms_stime field contains the
        CPU time spent in the system while executing tasks on behalf of the
        calling process." */
    printf("ticks: %lu\n", tms.tms_utime + tms.tms_stime);
    return 0;
}