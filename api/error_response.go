package api

import (
	"fmt"

	"github.com/hayashiki/mcfly-clone/api/apidata"
)

func NewApiErr(errorMessage string) *apidata.ApiError {
	return &apidata.ApiError{errorMessage}
}

func NewServerErr() *apidata.ApiError {
	return NewApiErr("Unknown server error. That's bad")
}

func NewAuthorizationHeaderRequiredErr() *apidata.ApiError {
	return NewApiErr("Authorization header required")
}

func NewInvalidAuthTokenError(token string) *apidata.ApiError {
	return NewApiErr(fmt.Sprintf("Auth token %s is not valid", token))
}

func NewInvalidTokenErr(token string) *apidata.ApiError {
	return NewApiErr(fmt.Sprintf("Invalid %s token", token))
}
