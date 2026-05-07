package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func login(username string, password string) bool {
	var b []byte

	// Convert to uppercase
	username = strings.ToUpper(username)

	// Strip everything after a space
	username, _, _ = strings.Cut(username, " ")

	b = system(
		"egrep \"^" + username + "\" /home/flag16/userdb.txt 2>&1",
	)

	for _, line := range strings.Split(string(b), "\n") {
		if _, pw, _ := strings.Cut(line, ":"); pw == password {
			return true
		}
	}

	return false
}

func main() {
	var mux *http.ServeMux = http.NewServeMux()
	var port int = 16000 + rand.Intn(50)
	var s *http.Server

	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, req *http.Request) {
			var ok bool = login(
				req.URL.Query().Get("username"),
				req.URL.Query().Get("password"),
			)
			var res string

			res = "<html>"
			res += "<head><title>Login results</title></head><body>"
			if ok {
				res += "Your login was accepted<br/>"
			} else {
				res += "Your login failed <br/>"
			}
			res += "Would you like a cookie?<br/><br/></body>"
			res += "</html>\n"

			w.Write([]byte(res))
		},
	)

	s = &http.Server{Addr: ":" + strconv.Itoa(port), Handler: mux}
	s.ListenAndServe()
}

func system(cmd string) []byte {
	var b []byte
	var c *exec.Cmd = exec.Command("/bin/bash", "-c", cmd)

	c.Env = []string{"PATH=" + os.Getenv("PATH")}
	b, _ = c.CombinedOutput()
	return bytes.TrimSpace(b)
}
