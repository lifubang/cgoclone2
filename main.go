package main

/*
#cgo CFLAGS: -Wall
#define _GNU_SOURCE
#include <endian.h>
#include <errno.h>
#include <fcntl.h>
#include <grp.h>
#include <limits.h>
#include <sched.h>
#include <setjmp.h>
#include <signal.h>
#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <unistd.h>

#include <sys/ioctl.h>
#include <sys/prctl.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/wait.h>

#include <linux/limits.h>
#include <linux/netlink.h>
#include <linux/types.h>

#include <sched.h>

#define STAGE_SETUP  -1
#define STAGE_PARENT  0
#define STAGE_CHILD   1
#define STAGE_INIT    2

int current_stage = STAGE_SETUP;

struct clone_t {
	char stack[4096] __attribute__((aligned(16)));
	char stack_ptr[0];

	jmp_buf *env;
	int jmpval;
};

static int child_func(void *arg) __attribute__((noinline));
static int child_func(void *arg)
{
	struct clone_t *ca = (struct clone_t *)arg;
	longjmp(*ca->env, ca->jmpval);
}

static int clone_parent(jmp_buf *env, int jmpval) __attribute__((noinline));
static int clone_parent(jmp_buf *env, int jmpval)
{
	struct clone_t ca = {
		.env = env,
		.jmpval = jmpval,
	};

	return clone(child_func, ca.stack_ptr, CLONE_PARENT | SIGCHLD, &ca);
}

static void nsexec() {
	jmp_buf env;
    current_stage = setjmp(env);
    switch (current_stage) {
		case STAGE_PARENT: {
			printf("STAGE_PARENT\n");
			clone_parent(&env, STAGE_CHILD);
			exit(0);
		}
		break;
		case STAGE_CHILD: {
			printf("STAGE_CHILD\n");
			clone_parent(&env, STAGE_INIT);
			exit(0);
		}
		break;
		case STAGE_INIT: {
			printf("STAGE_INIT\n");
		}
		break;
	}

	printf("This from nsexec\n");
	return;
}

void __attribute__((constructor)) init(void) {
	nsexec();
}
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println("From main!")
}
