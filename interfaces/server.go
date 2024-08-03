package interfaces

import (
	"github.com/private-project-pp/user-rpc-service/repository/postgre"
	"github.com/private-project-pp/user-rpc-service/shared/config"
	"github.com/private-project-pp/user-rpc-service/usecase/user"
)

func StartService() (err error) {
	config, err := config.SetupConfig()
	if err != nil {
		return err
	}

	db, err := SetupDatabase(config)
	if err != nil {
		return err
	}

	// setup repository
	userRepo := postgre.SetupUserRepo(db)

	// setup use case
	userService := user.SetupUserService(userRepo)

	_ = userService

	return nil
}
