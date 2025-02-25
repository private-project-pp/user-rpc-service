package interfaces

import (
	"fmt"
	"net"

	"github.com/private-project-pp/pos-general-lib/infrastructure"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/handler"
	"github.com/private-project-pp/user-rpc-service/repository/postgre"
	"github.com/private-project-pp/user-rpc-service/repository/redis"
	"github.com/private-project-pp/user-rpc-service/shared/config"
	"github.com/private-project-pp/user-rpc-service/usecase/authentication"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration"
	"google.golang.org/grpc/reflection"
)

func Container() (err error) {

	configs := config.SetupConfig()
	db, err := SetupDatabase(configs)
	if err != nil {
		return err
	}

	server := infrastructure.GrpcInstanceServer()
	reflection.Register(server)

	// setup repository
	authInfoRepo := postgre.SetupAuthInformationRepo(db)
	usersRepo := postgre.SetupUsersRepo(db)

	// Infrastructure repo
	redisRepo := redis.SetupRedis()

	// setup usecase
	authentication := authentication.SetupAuthService(authInfoRepo, redisRepo)
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
		return err
	}

	return nil
}
