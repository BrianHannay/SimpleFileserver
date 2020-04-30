package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	port := flag.Int("port", 8888, "Port number on which to listen")
	host := flag.String("host", "localhost", "Interface on which to listen")
	dir := flag.String("dir", ".", "Directory to serve")
	flag.Parse()
	realpath, err := filepath.Abs(*dir)
	if err != nil {
		panic(err)
	}
	addr := fmt.Sprintf(
		"%s:%d",
		*host,
		*port,
	)
	server := &http.Server{
		Addr:    addr,
		Handler: http.FileServer(http.Dir(realpath)),
	}
	fmt.Println("Serving", realpath, "on", addr)
	err = server.ListenAndServe()
	panic(err)
}
