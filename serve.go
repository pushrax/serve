package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	path := flag.String("path", ".", "Directory to serve")
	flag.Parse()

	fullPath, _ := filepath.Abs(*path)
	fullPath = filepath.Clean(fullPath)
	log.Println("Serving", fullPath, "on", *port)

	panic(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(fullPath))))
}
