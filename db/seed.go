package db

import (
	"github.com/hayashiki/mcfly-clone/models"
)

func (mdb *McflyDB) Seed() {
	mdb.Clean()

	u := &models.User{Name: strPtr("Matt Mockman"), AccessToken: "mock_seeded_access_token_123"}
	insertUser(mdb, u)
}

func insertUser(mdb *McflyDB, u *models.User) {
	err := mdb.SaveUser(u)
	if err != nil {
		return
	}
}

func strPtr(s string) *string {
	return &s
}
