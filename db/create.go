package db

import "fmt"

func (mdb *McflyDB) Create(databaseName string) {
	_, err := mdb.Exec(fmt.Sprintf("CREATE DATABASE %s", databaseName))
	if err != nil {
		return
	}
}
