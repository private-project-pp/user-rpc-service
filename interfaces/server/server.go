package server

import (
	"fmt"

	"github.com/private-project-pp/user-rpc-service/shared/config"
)

func StartService() (err error) {
	config := config.SetupConfig()

	fmt.Println(config.Service.Name)
	/*
		db, err := SetupDatabase(config)
		if err != nil {
			return err
		}

		// setup repository
		userRepo := postgre.SetupUserRepo(db)

		_ = userRepo*/

	return nil
}
