package main

// #include<dlfcn.h>
import "C"

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var b []byte
	var bin string
	var e error
	var lib string

	// TODO
	if bin, e = os.Executable(); e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	b, e = exec.Command("objdump", "-p", bin).CombinedOutput()
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "NEEDED") {
			lib = strings.TrimPrefix(line, "NEEDED")
			break
		}
	}

	lib = strings.TrimSpace(lib)
	C.dlopen(C.CString(lib), C.RTLD_NOW)

	fmt.Println("strace it!")
}
