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
	host := flag.String("host", "", "Host to listen on")
	flag.Parse()

	addr := *host + ":" + strconv.Itoa(*port)

	fullPath, _ := filepath.Abs(*path)
	fullPath = filepath.Clean(fullPath)

	log.Println("Serving", fullPath, "on", addr)

	panic(http.ListenAndServe(addr, http.FileServer(http.Dir(fullPath))))
}
