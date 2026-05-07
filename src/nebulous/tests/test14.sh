#!/usr/bin/env bash

cat >/tmp/flag14.go <<EOF
package main

import (
	"bufio"
	"fmt"
	"os"
)

func decrypt(s string, offset int) string {
	var b []rune = []rune(s)

	for i := range b {
		b[i] = rune(int(b[i]) - offset - i)
	}

	return string(b)
}

func main() {
	var f *os.File
	var n int
	var s *bufio.Scanner

    f, _ = os.Open(os.Args[1])
	for s = bufio.NewScanner(f); s.Scan(); {
		fmt.Println(decrypt(s.Text(), n))
		n += len(s.Text())
	}
}
EOF

flagpass="$(
    go run /tmp/flag14.go /home/flag14/passwd.txt | awk '{print $NF}'
)"

rm -f /tmp/flag14*

if [[ -n $flagpass ]]; then
    SSHPASS="$flagpass" sshpass -e -- \
        ssh -l flag14 -o StrictHostKeyChecking=no 127.0.0.1 getflag
fi
