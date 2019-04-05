package apidata

type LoginReq struct {
	Token    string `json:"token" validate:"nonzero"`
	Provider string `json:"provider" validate:"nonzero"`
}

type AipError struct {
	Error string `json:"error"`
}
