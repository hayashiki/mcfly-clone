package models

type Project struct {
	ID     int64  `db:"id" json:"id"`
	Handle string `db:"handle" json:"handle"`
	// TODO : fix this and test
	SourceUrl      string `db:"source_url" json:"username"`
	SourceProvider string `db:"source_provider" json:"source_provider"`
}

// func (db *DB)
