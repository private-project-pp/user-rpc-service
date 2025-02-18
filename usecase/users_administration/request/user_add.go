package request

type UserAddRequest struct {
	Fullname    string `json:""`
	Email       string `json:""`
	PhoneNumber string `json:""`
}

func NewUserAddRequest(in interface{}) (out UserAddRequest, err error) {
	return out, nil
}
