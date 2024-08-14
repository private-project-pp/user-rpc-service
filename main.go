package main

import (
	"github.com/private-project-pp/user-rpc-service/interfaces/server"
	_ "github.com/spf13/viper"
	_ "google.golang.org/grpc"
	_ "gorm.io/gorm"
)

func main() {
	err := server.StartServer()
	if err != nil {
		panic(err)
	}
}
