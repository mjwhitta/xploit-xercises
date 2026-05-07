package main

import (
	"bufio"
	"bytes"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	var a *net.TCPAddr
	var b []byte
	var c *net.TCPConn
	var e error
	var in string
	var l *net.TCPListener
	var port int = 12000 + rand.Intn(50)

	a, _ = net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	if a != nil {
		l, _ = net.ListenTCP("tcp", a)
	}

	for l != nil {
		if c, _ = l.AcceptTCP(); c == nil {
			continue
		}

		c.Write([]byte("Password: "))
		c.SetReadDeadline(time.Now().Add(time.Minute))

		if in, e = bufio.NewReader(c).ReadString('\n'); e == nil {
			in = strings.TrimSpace(in)
			b = system("echo -n " + in + " | sha1sum")

			if string(b[0:8]) != "025abc01" {
				c.Write([]byte("Better luck next time\n"))
			} else {
				c.Write([]byte("Congrats, your token is REDACTED!\n"))
			}
		}

		c.Close()
	}
}

func system(cmd string) []byte {
	var b []byte
	var c *exec.Cmd = exec.Command("/bin/bash", "-c", cmd)

	c.Env = []string{"PATH=" + os.Getenv("PATH")}
	b, _ = c.CombinedOutput()
	return bytes.TrimSpace(b)
}
