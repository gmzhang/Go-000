package errs

import (
	"errors"
)

const ErrCodeOK = 1000
const ErrCodeSegmentBase = 4000

var (
	ErrInvalidParam = errors.New("参数错误")
	ErrUserNotFound = errors.New("用户不存在")
)

func GetErrorCode(err error) int32 {
	switch err {
	case nil:
		return ErrCodeOK
	case ErrInvalidParam:
		return ErrCodeSegmentBase + 1
	case ErrUserNotFound:
		return ErrCodeSegmentBase + 2

	default:
		return -1
	}
}

func GetErrorMap(err error) map[string]interface{} {
	var msg = "OK"
	if err != nil {
		msg = err.Error()
	}

	return map[string]interface{}{
		"errCode": GetErrorCode(err),
		"msg":     msg,
	}
}
