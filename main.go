package main

import (
	"flag"
	"net/http"
)

func main() {
	const (
		defaultPort = "8080"
		portUsage   = "Port to serve on"
		defaultFile = "index.html"
		fileUsage   = "Static file to serve"
	)

	var (
		port string
		file string
	)

	flag.StringVar(&port, "port", defaultPort, portUsage)
	flag.StringVar(&port, "p", defaultPort, portUsage+" (shorthand)")
	flag.StringVar(&file, "file", defaultFile, fileUsage)
	flag.StringVar(&file, "f", defaultFile, fileUsage+" (shorthand)")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	})

	panic(http.ListenAndServe(":"+port, nil))
}
