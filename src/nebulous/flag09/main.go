package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var cmd []string = []string{"php", "/home/flag09/.flag09.php"}

	syscall.Setresgid(os.Getegid(), os.Getegid(), os.Getegid())
	syscall.Setresuid(os.Geteuid(), os.Geteuid(), os.Geteuid())

	system(append(cmd, os.Args[1:]...))
}

func system(cmd []string) {
	var c *exec.Cmd = exec.Command(cmd[0], cmd[1:]...)

	c.Env = []string{"PATH=" + os.Getenv("PATH")}
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Run()
}
