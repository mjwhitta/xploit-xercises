package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	var e error
	var u *user.User

	if u, e = user.Current(); e != nil {
		fmt.Println("couldn't find account")
		os.Exit(1)
	}

	if u.Username == "root" {
		fmt.Printf(
			"\x1b[33m[-] %s\x1b[0m\n",
			"While you got root, that wasn't the intended target -_-",
		)
	} else if strings.HasPrefix(u.Username, "flag") {
		fmt.Printf(
			"\x1b[32m[+] %s %s %s\x1b[0m\n",
			"You have successfully executed getflag on target",
			"account",
			u.Username,
		)
	} else {
		fmt.Printf(
			"\x1b[31m[!] %s %s\x1b[0m\n",
			"getflag is executing on a non-flag account, this",
			"doesn't count",
		)
	}
}
