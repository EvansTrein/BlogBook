package main

import "backend/server"

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
}
