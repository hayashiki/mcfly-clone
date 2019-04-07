package models

type User struct {
	ID          int64   `db:"id" json:"id"`
	Name        *string `db:"name" json:"name"`
	AccessToken string  `db:"access_token" json:"access_token`
}

type ProviderAccessToken struct {
	Provider         string `db:"provider" json:"provider"`
	ProviderUsername string `db:"provider_username" json:"provider_username"`
	AccessToken      string `db:"access_token" json:"access_token"`
	UserID           int64  `db:"user_id" json:"user_id"`
}

func (db *DB) SaveUser(u *User) error {
	var id int64

	q := `INSERT INTO mcfly_user(name, access_token) VALUES(s1, s2) RETURNING id`
	err := db.QueryRowScan(&id, q, u.Name, u.AccessToken)
	if err != nil {
		return nil
	}
	u.ID = id
	return nil
}
