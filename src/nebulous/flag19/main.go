package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var e error
	var fi os.FileInfo
	var proc string = fmt.Sprintf("/proc/%d", os.Getppid())

	// Stat the parent's /proc entry
	if fi, e = os.Stat(proc); e != nil {
		fmt.Println("unable to check parent process")
		os.Exit(1)
	}

	// Check the owner ID
	if fi.Sys().(*syscall.Stat_t).Uid == 0 {
		// If root, it is ok to start a shell
		syscall.Setresgid(os.Getegid(), os.Getegid(), os.Getegid())
		syscall.Setresuid(os.Geteuid(), os.Geteuid(), os.Geteuid())
		system(os.Args[1:])
		os.Exit(0)
	}

	fmt.Println("You are unauthorized to run this program")
}

func system(cmd []string) {
	var c *exec.Cmd = exec.Command(cmd[0], cmd[1:]...)

	c.Env = []string{"PATH=" + os.Getenv("PATH")}
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Run()
}
