package api_test

import (
	"github.com/hayashiki/mcfly-clone/client"
	"github.com/hayashiki/mcfly-clone/db"
)

var (
	apiClient *client.McflyClient
)

func resetDB() {
	cleanupDB()
	seedDB()
}

func cleanupDB() {
	mdb.Clean()
}

func seedDB() {
	mdb.Seed()
}