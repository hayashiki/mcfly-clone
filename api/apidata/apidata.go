package apidata

type LoginReq struct {
	Token    string `json:"token" validate:"nonzero"`
	Provider string `json:"provider" validate:"nonzero"`
}

type ApiResponse struct {
	Status string `json:"status"`
}

type ApiError struct {
	Error string `json:"error"`
}

type ProjectReq struct {
	Handle   string `json:"handle" validate:"nonzero"`
	Provider string `json:"provider" validate:"nonzero"`
}

type ProjectResp struct {
	Handle   string `json:"handle"`
	Url      string `json:"url"`
	Provider string `json:"provider"`
}
