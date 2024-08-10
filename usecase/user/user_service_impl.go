package user

type userService struct {
}

func SetupUserService() UserServiceServer {
	return &userService{}
}
