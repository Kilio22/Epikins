package util

import (
	"epikins-api/internal"
)

func GetMyError(message string, err error, code int) internal.MyError {
	if err != nil {
		return internal.MyError{
			Message: message + ": " + err.Error(),
			Status:  code,
		}
	}
	return internal.MyError{
		Message: message,
		Status:  code,
	}
}
