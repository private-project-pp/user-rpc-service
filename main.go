package main

import (
	"github.com/private-project-pp/user-rpc-service/interfaces/server"
)

func main() {
	err := server.StartServer()
	if err != nil {
		panic(err)
	}
}
