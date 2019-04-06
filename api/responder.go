package api

import (
	"encoding/json"
	"net/http"

	"github.com/hayashiki/mcfly-clone/api/apidata"
)

type Responder struct {
	http.ResponseWriter
	Request *http.Request
}

func (r *Responder) Respond(v interface{}) {
	r.WriteCommonHeaders()
	r.WriteSuccessHeaders()
	r.WriteResponseData(v)
}

func (r *Responder) RespondWithError(apiErr *apidata.ApiError) {
	r.WriteCommonHeaders()
	r.WriteErrorHeaders()
	r.WriteResponseData(apiErr)
}

func (r *Responder) WriteCommonHeaders() {
	r.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func (r *Responder) WriteErrorHeaders() {
	r.WriteHeader(http.StatusBadRequest)
}

func (r *Responder) WriteSuccessHeaders() {
	r.WriteHeader(http.StatusOK)
}

func (r *Responder) WriteResponseData(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
	}
	r.Write(b)
}

func (r *Responder) RespondWithServerError(err error) {
	r.WriteCommonHeaders()
	r.WriteErrorHeaders()
	r.WriteResponseData(NewServerErr())
}

func (r *Responder) RespondWithSuccess() {
	r.Respond(NewSuccessResponse())
}

func NewSuccessResponse() *apidata.ApiResponse {
	return &apidata.ApiResponse{"success!!"}
}
