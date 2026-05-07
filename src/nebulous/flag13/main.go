package main

// #include<unistd.h>
import "C"

import (
	"fmt"
	"os"
)

const wantUID C.uint = 9001

func main() {
	if C.getuid() != wantUID {
		fmt.Printf(
			"Security failure detected. Got UID %d, want %d.\n",
			C.getuid(),
			wantUID,
		)
		fmt.Println("System administrators will be notified.")
		os.Exit(1)
	}

	fmt.Println("FLAGPASS")
}
