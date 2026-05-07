package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	var mux *http.ServeMux = http.NewServeMux()
	var port int = 7000 + rand.Intn(50)
	var s *http.Server

	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, req *http.Request) {
			var b []byte
			var host string = req.URL.Query().Get("host")
			var res string

			b = system("ping -c 1 " + host)

			res = "<html>"
			res += "<head><title>Ping results</title></head>"
			res += "<body><pre>" + string(b) + "</pre></body>"
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
