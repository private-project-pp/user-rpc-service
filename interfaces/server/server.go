package server

import (
	"fmt"
	"net"

	"github.com/private-project-pp/user-rpc-service/interfaces"
	"github.com/private-project-pp/user-rpc-service/model"
	"github.com/private-project-pp/user-rpc-service/repository/postgre"
	"github.com/private-project-pp/user-rpc-service/shared/config"
	"github.com/private-project-pp/user-rpc-service/usecase/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer() (err error) {

	configs := config.SetupConfig()
	db, err := interfaces.SetupDatabase(configs)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	reflection.Register(server)

	// setup repository
	authInfoRepo := postgre.SetupAuthInformationRepo(db)

	// setup service-usecase
	authService := authentication.SetupAuthService(authInfoRepo)

	model.RegisterAuthenticationServiceServer(server, authService)

	listen, err := net.Listen("tcp", configs.Service.Port)
	if err != nil {
		return err
	}

	fmt.Println("RUNNING GRPC")
	fmt.Printf("Running on PORT [%s]", config.Service.Port)
	if err = server.Serve(listen); err != nil {
		panic(err)
	}

	return nil
}
