package server

import (
	"github.com/private-project-pp/user-rpc-service/interfaces"
	"github.com/private-project-pp/user-rpc-service/repository/postgre"
	"github.com/private-project-pp/user-rpc-service/shared/config"
)

func StartService() (err error) {
	config := config.SetupConfig()
	db, err := interfaces.SetupDatabase(config)
	if err != nil {
		return err
	}

	// setup repository
	authInfoRepo := postgre.SetupAuthInformationRepo(db)

	_ = authInfoRepo
	return nil
}
