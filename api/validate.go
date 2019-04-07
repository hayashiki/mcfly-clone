package api

import (
	"github.com/hayashiki/mcfly-clone/models"
)

func (r *Responder) ValidateAuthorization(db *models.Datastore) *models.User {
	authHeader := r.Request.Header["Authorization"]
	if len(authHeader) == 0 {
		return nil
	}
}
