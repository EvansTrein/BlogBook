package main

import (
	"authForum/logging"
	"authForum/server"
)

func init() {
	logging.InitLogger()
	server.InitServer()
}

func main() {
	server.StartServer()
}
