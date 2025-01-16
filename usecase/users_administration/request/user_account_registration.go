package request

import "encoding/json"

type UserAccountRegistrationReq struct {
	UserId   string `json:"userId" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewUserAccountRegistrationReq(in []byte) (out UserAccountRegistrationReq, err error) {
	if err = json.Unmarshal(in, &out); err != nil {
		return out, err
	}
	return out, nil
}
