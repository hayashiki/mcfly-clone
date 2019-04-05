package api

import (
	"github.com/hayashiki/mcfly-clone/models"
)

type RequestContext struct {
	CurrentUser *models.User
}
