#!/usr/bin/env python3

import os
import pickle
import random
import signal
import socket
import time

signal.signal(signal.SIGCHLD, signal.SIG_IGN)

def handle(clnt):
    line = clnt.recv(1024)
    pickle.loads(line)
    clnt.close()

skt = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
skt.bind(("0.0.0.0", 17000 + random.randrange(0, 50)))
skt.listen(10)

while True:
    clnt, addr = skt.accept()
    if os.fork() == 0:
        handle(clnt)
        exit(1)
