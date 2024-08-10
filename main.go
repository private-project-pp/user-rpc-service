package main

import (
	"fmt"

	"github.com/private-project-pp/user-rpc-service/interfaces/server"
	"github.com/private-project-pp/user-rpc-service/shared/config"
	_ "github.com/spf13/viper"
	_ "google.golang.org/grpc"
	_ "gorm.io/gorm"
)

func main() {
	server.StartService()
	fmt.Println("Config from main =>", config.Service.Name)
}
