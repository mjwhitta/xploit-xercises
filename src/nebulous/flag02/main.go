package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var cmd string
	var s string

	syscall.Setresgid(os.Getegid(), os.Getegid(), os.Getegid())
	syscall.Setresuid(os.Geteuid(), os.Geteuid(), os.Geteuid())

	cmd = fmt.Sprintf("/bin/echo %s is cool", os.Getenv("USER"))
	s = "/bin/echo $USER is cool"
	fmt.Printf("about to call system(\"%s\")\n", s)
	system(cmd)
}

func system(cmd string) {
	var c *exec.Cmd = exec.Command("/bin/bash", "-c", cmd)

	c.Env = []string{"PATH=" + os.Getenv("PATH")}
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Run()
}
