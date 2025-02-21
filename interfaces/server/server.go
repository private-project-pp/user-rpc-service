package server

import (
	"fmt"
	"net"

	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/handler"
	"github.com/private-project-pp/user-rpc-service/interfaces"
	"github.com/private-project-pp/user-rpc-service/repository/postgre"
	"github.com/private-project-pp/user-rpc-service/shared/config"
	"github.com/private-project-pp/user-rpc-service/usecase/authentication"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration"

	"github.com/private-project-pp/pos-general-lib/infrastructure"
	"google.golang.org/grpc/reflection"
)

func StartServer() (err error) {

	configs := config.SetupConfig()
	db, err := interfaces.SetupDatabase(configs)
	if err != nil {
		return err
	}

	server := infrastructure.GrpcInstanceServer()
	reflection.Register(server)

	// setup repository
	authInfoRepo := postgre.SetupAuthInformationRepo(db)
	usersRepo := postgre.SetupUsersRepo(db)

	// setup usecase
	authentication := authentication.SetupAuthService(authInfoRepo)
	userAdministration := users_administration.SetupUserAdministration(authInfoRepo, usersRepo)

	//setup RPC handler
	rpcHandler := handler.SetupUserService(authentication, userAdministration)

	model.RegisterUserServiceServer(server, rpcHandler)

	// setup service-usecase
	// authService := authentication.SetupAuthService(authInfoRepo)
	// userAdmService := users_administration.SetupUserAdministration(authInfoRepo, usersRepo)
	// model.RegisterAuthenticationServiceServer(server, authService)
	// model.RegisterUsersAdministrationServiceServer(server, userAdmService)

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
