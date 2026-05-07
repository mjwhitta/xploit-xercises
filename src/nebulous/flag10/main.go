package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"syscall"
)

func main() {
	var a *net.TCPAddr
	var b []byte
	var c *net.TCPConn
	var e error
	var file string
	var host string

	if len(os.Args) != 3 {
		fmt.Println("flag10 <file> <host>")
		os.Exit(1)
	}

	file = os.Args[1]
	host = os.Args[2]

	if e = syscall.Access(file, 0x4); e != nil { // R_OK == 0x4
		fmt.Printf("you do not have access to %s\n", file)
		os.Exit(1)
	}

	fmt.Printf("Connecting to %s:18211...\n", host)

	if a, e = net.ResolveTCPAddr("tcp", host+":18211"); e != nil {
		fmt.Printf("unable to resolve host %s\n", host)
		os.Exit(1)
	}

	if c, e = net.DialTCP("tcp", nil, a); e != nil {
		fmt.Printf("unable to connect to host %s\n", host)
		os.Exit(1)
	}

	if _, e = c.Write([]byte(".oO Oo.\n")); e != nil {
		fmt.Printf("unable to write banner to host %s\n", host)
		os.Exit(1)
	}

	fmt.Printf("Connected!\nSending file...\n")

	if b, e = os.ReadFile(file); e != nil {
		fmt.Printf("unable to read from file %s\n", file)
		os.Exit(1)
	}

	if _, e = c.Write(append(bytes.TrimSpace(b), '\n')); e != nil {
		fmt.Printf("unable to write file to host %s\n", host)
		os.Exit(1)
	}

	fmt.Println("Wrote file!")
}
