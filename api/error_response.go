package api

import (
	"github.com/hayashiki/mcfly-clone/api/apidata"
)

func NewApiErr(errorMessage string) *apidata.ApiError {
	return &apidata.ApiError{errorMessage}
}

func NewServerErr() *apidata.ApiError {
	return NewApiErr("Unknown server error. That's bad")
}
