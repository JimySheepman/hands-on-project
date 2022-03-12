/* -*- compile-command: "gcc -Wall -Werror setfcap_and_exec.c -o setfcap_and_exec  -static -lcap" -*- */
#include <errno.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <linux/capability.h>
#include <sys/capability.h>
#include <sys/prctl.h>
#include <sys/types.h>

int main (int argc, char  **argv)
{
	if (argc == 2 && !strcmp(argv[1], "inner")) {
		cap_t self_caps = {0};
		if (!(self_caps = cap_get_proc())) {
			fprintf(stderr, "++ cap_get_proc failed: %m\n");
			return 1;
		}

		cap_flag_value_t cap_mknod_status = CAP_CLEAR;
		if (cap_get_flag(self_caps, CAP_MKNOD, CAP_PERMITTED, &cap_mknod_status)) {
			fprintf(stderr, "++ cap_get_flag failed: %m\n");
			cap_free(self_caps);
			return 1;
		}
		if (cap_mknod_status == CAP_CLEAR)
			fprintf(stderr, "!! don't have cap_mknod+p?\n");

		if (cap_set_flag(self_caps, CAP_EFFECTIVE, 1,
				 & (cap_value_t) { CAP_MKNOD }, CAP_SET)) {
			fprintf(stderr, "++ can't cap_set_flag: %m\n");
			cap_free(self_caps);
			return 1;
		}
		if (cap_set_proc(self_caps)) {
			fprintf(stderr, "++ can't cap_set_proc: %m\n");
			cap_free(self_caps);
			return 1;
		}
		cap_free(self_caps);
		fprintf(stderr, "++ have CAP_MKNOD!\n");
	} else {
		cap_t file_caps = {0};
		if (!(file_caps = cap_from_text("cap_mknod+p"))) {
			fprintf(stderr, "++ cap_from_text failed: %m\n");
			return 1;
		}
		if (cap_set_file(argv[0], file_caps)) {
			fprintf(stderr, "++ cap_set_file failed: %m\n");
			cap_free(file_caps);
			return 1;
		}
		cap_free(file_caps);

		if (execve(argv[0], (char  *[]){ argv[0], "inner", 0 }, NULL)) {
			fprintf(stderr, "++ execve failed: %m\n");
			return 1;
		}
	}
	return 0;
}
