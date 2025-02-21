package request

import (
	"encoding/json"

	"github.com/private-project-pp/pos-general-lib/stacktrace"
)

type UserAccountRegistrationReq struct {
	UserId   string `json:"userId" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewUserAccountRegistrationReq(in any) (out UserAccountRegistrationReq, err error) {
	byteData, err := json.Marshal(in)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}
	if err = json.Unmarshal(byteData, &out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}
	return out, nil
}
