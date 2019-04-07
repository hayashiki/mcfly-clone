package api

import (
	"net/http"

	"github.com/hayashiki/mcfly-clone/models"
	"github.com/hayashiki/mcfly-clone/provider"
)

type Handlers struct {
	DB            models.Datastore
	GenerateToken func() string
}

type HandlerOptions struct {
	AuthRequired       bool
	RequestData        interface{}
	ProjectContext     bool
	UserSourceProvider bool
	UserAuthProvider   bool
	After              func(*Responder, *RequestContext)
}

type RequestContext struct {
	CurrentUser    *models.User
	RequestData    interface{}
	SourceProvider *provider.SourceProvider
	AuthProvider   *provider.AuthProvider
}

func (handlers *Handlers) MakeHandlerFunc(opts HandlerOptions) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
