package main

import (
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("/usr/share/files"))

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	http.ListenAndServe(":"+httpPort, fs)
}
