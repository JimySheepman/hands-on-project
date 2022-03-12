/* -*- compile-command: "gcc -Wall -Werror -static ptrace_breaks_seccomp.c -o ptrace_breaks_seccomp" -*- */
#include <sys/stat.h>
#include <stdio.h>
#include <sys/ptrace.h>
#include <unistd.h>
#include <sys/types.h>
#include <signal.h>
#include <sys/user.h>
#include <sys/wait.h>
#include <stddef.h>
#include <sys/syscall.h>

#define MAGIC_SYSCALL 666

int main(int argc, char **argv)
{
    pid_t child = 0;
    switch ((child = fork()))
    {
    case -1:
        fprintf(stderr, "++ fork failed: %m\n");
        return 1;
    case 0:;
        fprintf(stderr, "++ child stopping itself.\n");
        if (kill(getpid(), SIGSTOP))
        {
            fprintf(stderr, "++ kill failed: %m\n");
            return 1;
        }
        fprintf(stderr, "++ child continued\n");
        /* pick an arbitrary syscall number. our tracer will change it to chmod. */
        if (syscall(MAGIC_SYSCALL, argv[0], S_ISUID | S_IRUSR | S_IWUSR | S_IXUSR))
        {
            fprintf(stderr, "chmod-via-nanosleep failed: %m\n");
            return 1;
        }
        fprintf(stderr, "++ chmod succeeded, child finished.\n");
        break;
    default:;
        int status = 0;
        if (ptrace(PTRACE_ATTACH, child, NULL, NULL))
        {
            fprintf(stderr, "++ ptrace failed: %m\n");
            return 1;
        }
        waitpid(child, &status, 0);
        if (!(status & SIGSTOP))
        {
            fprintf(stderr, "++ expected SIGSTOP in child.\n");
            return 1;
        }
        struct user_regs_struct regs = {0};
        while (1)
        {
            if (ptrace(PTRACE_GETREGS, child, 0, &regs))
            {
                fprintf(stderr, "++ getting child registers failed: %m\n");
                return 1;
            }
            if (!(regs.orig_rax == MAGIC_SYSCALL))
            {
                if (ptrace(PTRACE_SYSCALL, child, 0, 0))
                {
                    fprintf(stderr, "++ continuing the process failed.\n");
                    return 1;
                }
                waitpid(child, &status, 0);
                if (!(status & SIGTRAP))
                {
                    fprintf(stderr, "++ expected SIGTRAP in child.\n");
                    return 1;
                }
            }
            else
            {
                fprintf(stderr, "++ got MAGIC_SYSCALL!\n");
                regs.orig_rax = SYS_chmod;
                if (ptrace(PTRACE_SETREGS, child, 0, &regs))
                {
                    fprintf(stderr, "++ continuing child failed: %m\n");
                    return 1;
                }
                if (ptrace(PTRACE_CONT, child, 0, 0))
                {
                    fprintf(stderr, "++ continuing child failed: %m\n");
                    return 1;
                }
                break;
            }
        }
        waitpid(child, NULL, 0);
        fprintf(stderr, "++ finished waiting.\n");

        break;
    }
    return 0;
}
