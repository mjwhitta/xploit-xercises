package main

import (
	"bufio"
	"fmt"
	"os"
)

func encrypt(s string, offset int) string {
	var b []rune = []rune(s)

	for i := range b {
		b[i] = rune(int(b[i]) + offset + i)
	}

	return string(b)
}

func main() {
	var e error
	var f *os.File = os.Stdin
	var n int
	var s *bufio.Scanner

	if len(os.Args) > 1 {
		if f, e = os.Open(os.Args[1]); e != nil {
			fmt.Println(e.Error())
			os.Exit(1)
		}
	}

	for s = bufio.NewScanner(f); s.Scan(); {
		fmt.Println(encrypt(s.Text(), n))
		n += len(s.Text())
	}

	if s.Err() != nil {
		fmt.Println(s.Err().Error())
		os.Exit(1)
	}
}
