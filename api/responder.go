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

func (r *Responder) RespondWithError(apiErr *apidata.AipError) {
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

func (r *Responder) WriteResponseData(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
	}
	r.Write(b)
}
