package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var b []byte
	var e error

	if len(os.Args) != 2 {
		fmt.Println("flag04 [file to read]")
		os.Exit(1)
	}

	if strings.Contains(os.Args[1], "passwd") {
		fmt.Printf("you may not access '%s'\n", os.Args[1])
		os.Exit(1)
	}

	if b, e = os.ReadFile(os.Args[1]); e != nil {
		fmt.Printf("unable to read %s\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Println(string(bytes.TrimSpace(b)))
}
