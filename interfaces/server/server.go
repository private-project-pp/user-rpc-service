package server

import (
	"fmt"
	"net"

	"github.com/private-project-pp/user-rpc-service/interfaces"
	"github.com/private-project-pp/user-rpc-service/repository/postgre"
	"github.com/private-project-pp/user-rpc-service/shared/config"
	"github.com/private-project-pp/user-rpc-service/usecase/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer() (err error) {

	configs := config.SetupConfig()
	db, err := interfaces.SetupDatabase(configs)
	if err != nil {
		return err
	}

	// setup repository
	authInfoRepo := postgre.SetupAuthInformationRepo(db)

	userService := user.SetupUserService(authInfoRepo)

	server := grpc.NewServer()
	reflection.Register(server)

	user.RegisterUserServiceServer(server, userService)

	listen, err := net.Listen("tcp", configs.Service.Port)
	if err != nil {
		return err
	}

	fmt.Println("RUNNING GRPC")
	if err = server.Serve(listen); err != nil {
		panic(err)
	}

	return nil
}
