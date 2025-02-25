package responses

type UserListResponse struct {
	Users []UserDataResponse
}

type UserDataResponse struct {
	Name string
}
