package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	syscall.Setresgid(os.Getegid(), os.Getegid(), os.Getegid())
	syscall.Setresuid(os.Geteuid(), os.Geteuid(), os.Geteuid())
	system("/bin/bash -l")
}

func system(cmd string) {
	var c *exec.Cmd = exec.Command("/bin/bash", "-c", cmd)

	c.Env = []string{"PATH=" + os.Getenv("PATH")}
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Run()
}
