#!/usr/bin/env bash

for i in /home/flag03/writable.d/*; do
    bash "$i"
    rm -f "$i"
done; unset i
