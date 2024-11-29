package main

import "authForum/server"

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
}
