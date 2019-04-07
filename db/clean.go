package db

import _ "github.com/lib/pq"

func (mdb *McflyDB) Clean() {
	tx, err := mdb.Begin()
	if err != nil {
		return
	}
	tx.Exec("DELETE FROM build")
	tx.Exec("DELETE FROM user_project")
	tx.Exec("DELETE FROM project")
	tx.Exec("DELETE FROM provider_access_token")
	tx.Exec("DELETE FROM mcfly_user")
	tx.Commit()
}
