package main

import "api/server"

func main() {
	s := server.NewHTTPServer("8080")
	s.Serve()
}
