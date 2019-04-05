package models

type User struct {
	ID          int64   `db:"id" json:"id"`
	Name        *string `db:"name" json:"name"`
	AccessToken string  `db:"access_token" json:"access_token`
}
