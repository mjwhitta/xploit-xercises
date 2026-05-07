package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Return a random, non-predictable file, and return the file pointer
// for it.
func getRand() *os.File {
	var e error
	var f *os.File
	var path string
	var pid int = os.Getpid()
	var r *rand.Rand
	var tmp string = os.Getenv("TEMP")

	r = rand.New(rand.NewSource(time.Now().Unix()))

	path = fmt.Sprintf(
		"%s/%d.%c%c%c%c%c%c",
		tmp,
		pid,
		'A'+r.Intn(26),
		'0'+r.Intn(10),
		'a'+r.Intn(26),
		'A'+r.Intn(26),
		'0'+r.Intn(10),
		'a'+r.Intn(26),
	)

	f, e = os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0o600)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	return f
}

func main() {
	var blue int
	var buf [1024]byte
	var cl string = "Content-Length: "
	var e error
	var f *os.File
	var length int
	var line string
	var mem []byte
	var pink int
	var fgets *bufio.Reader = bufio.NewReader(os.Stdin)

	// The original C code didn't have this, but if you start by
	// telling people there are two solutions, there better be two
	// solutions.
	syscall.Setresgid(os.Getegid(), os.Getegid(), os.Getegid())
	syscall.Setresuid(os.Geteuid(), os.Geteuid(), os.Geteuid())

	if line, e = fgets.ReadString('\n'); e != nil {
		fmt.Println("reading from stdin")
		os.Exit(1)
	}

	line = strings.TrimSuffix(line, "\n")
	if len(line) > 256 {
		line = line[:256]
	}

	if !strings.HasPrefix(line, cl) {
		fmt.Println("invalid header")
		os.Exit(1)
	}

	length, _ = strconv.Atoi(strings.TrimPrefix(line, cl))

	if length < len(buf) {
		for i := 0; i < length; i++ {
			if buf[i], e = fgets.ReadByte(); e != nil {
				fmt.Println("fread length")
				os.Exit(1)
			}
		}

		// If C, this buf would contain uninitialized garbage. Go
		// doesn't have that problem. Ultimately this doesn't affect
		// the challenge as you could just execute the C program
		// repeatedly until a NULL byte shows up in the right place.
		process(buf[:length], length)
	} else {
		f = getRand()
		defer f.Close()

		for blue = length; blue > 0; {
			pink = len(buf)
			for i := 0; i < len(buf); i++ {
				if buf[i], e = fgets.ReadByte(); e != nil {
					pink = i
					break
				}
			}

			fmt.Printf(
				"blue = %d, length = %d, pink = %d\n",
				blue,
				length,
				pink,
			)

			if pink == 0 {
				fmt.Printf(
					"fread fail(blue = %d, length = %d)\n",
					blue,
					length,
				)
				os.Exit(1)
			}

			f.Write(buf[:pink])
			blue -= pink
		}

		mem, e = syscall.Mmap(
			int(f.Fd()),
			0,
			length,
			syscall.PROT_READ|syscall.PROT_WRITE,
			syscall.MAP_PRIVATE,
		)
		if e != nil {
			fmt.Println("mmap")
			os.Exit(1)
		}

		process(mem[:length], length)
	}
}

func process(buf []byte, length int) {
	var key uint8 = uint8(length) & 0xff

	for i := 0; i < length; i++ {
		buf[i] = buf[i] ^ key
		key -= buf[i]
	}

	system(string(buf))
}

func system(cmd string) {
	var c *exec.Cmd = exec.Command("/bin/bash", "-c", cmd)

	c.Env = []string{"PATH=" + os.Getenv("PATH")}
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Run()
}
