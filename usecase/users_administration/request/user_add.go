package request

import (
	"encoding/json"

	"github.com/private-project-pp/pos-general-lib/stacktrace"
)

type UserAddRequest struct {
	Fullname    string `json:"fullname" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

func NewUserAddRequest(in interface{}) (out UserAddRequest, err error) {
	byteData, err := json.Marshal(in)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	err = json.Unmarshal(byteData, &out)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}
