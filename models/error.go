package models

import (
	"errors"
)

var ErrNotFound = errors.New("note found")
var ErrExists = errors.New("already exists")
var ErrDuplicate = errors.New("duplicate")
var ErrUnknown = errors.New("unknown model error")

type QueryExecError struct {
	Query   string
	DbError error
	Name    string
	Arges   []interface{}
	NoRows  bool
	TxDone  bool
}
