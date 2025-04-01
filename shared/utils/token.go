package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/private-project-pp/pos-general-lib/shared/constant"
	utils_lib "github.com/private-project-pp/pos-general-lib/shared/utils"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	"github.com/private-project-pp/user-rpc-service/shared/config"
)

type JwtHeader struct {
	Alg constant.HashType `json:"alg"`
}

type JwtPayload struct {
	XAuthUserId string `json:"X-AuthUser-Id"`
	RoleCode    string `json:"X-UserRole-Code"`
	CreatedTime string `json:"X-CreatedTime"`
}

func GenerateUserToken(alg constant.HashType, sessionInfo map[string]string) (out string, err error) {
	jwtHead := JwtHeader{
		Alg: alg,
	}

	byteData, err := json.Marshal(jwtHead)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	jwtHeadBase64 := base64.RawStdEncoding.EncodeToString(byteData)

	byteData, err = json.Marshal(sessionInfo)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	jwtPayloadBase64 := base64.RawStdEncoding.EncodeToString(byteData)

	mergedhash := fmt.Sprintf("%s.%s", jwtHeadBase64, jwtPayloadBase64)
	out, err = utils_lib.GenerateSHA256Hash(mergedhash, config.Auth.JwtSecureKey)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.BAD_PROCESSING, err.Error())
	}

	out = fmt.Sprintf("%s.%s", mergedhash, out)

	return out, nil
}
