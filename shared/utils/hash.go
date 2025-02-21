package utils

import (
	"github.com/private-project-pp/pos-general-lib/shared/utils"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
)

func GeneratePasswordHash(value, key string) (hashed string, err error) {

	hashed, err = utils.Generate256Hash(value, key)
	if err != nil {
		return hashed, stacktrace.Cascade(err, stacktrace.BAD_PROCESSING, err.Error())
	}
	return hashed, nil
}
