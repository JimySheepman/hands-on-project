/* Local Variables: */
/* compile-command: "gcc -Wall -Werror -static  subverting_networking.c \*/
/*                   -o subverting_networking" */
/* End: */
#define _GNU_SOURCE
#include <stdio.h>
#include <unistd.h>
#include <sched.h>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <linux/sockios.h>

int main (int argc, char **argv)
{
	if (unshare(CLONE_NEWUSER | CLONE_NEWNET)) {
		fprintf(stderr, "++ unshare failed: %m\n");
		return 1;
	}
	/* this is how you create a bridge... */
	int sock = 0;
	if ((sock = socket(PF_LOCAL, SOCK_STREAM, 0)) == -1) {
		fprintf(stderr, "++ socket failed: %m\n");
		return 1;
	}
	if (ioctl(sock, SIOCBRADDBR, "br0")) {
		fprintf(stderr, "++ ioctl failed: %m\n");
		close(sock);
		return 1;
	}
	close(sock);
	fprintf(stderr, "++ success!\n");
	return 0;
}
