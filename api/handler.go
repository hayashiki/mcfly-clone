package api

import (
	"github.com/hayashiki/mcfly-clone/models"
	"github.com/hayashiki/mcfly-clone/provider"
)

type Handlers struct {
	DB models.Datastore
}

type RequestContext struct {
	CurrentUser    *models.User
	RequestData    interface{}
	SourceProvider *provider.SourceProvider
	AuthProvider   *provider.AuthProvider
}
