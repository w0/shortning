package main

import "github.com/w0/shortning/internal/server"

func main() {
	server := server.NewServer()
	server.ListenAndServe()
}
