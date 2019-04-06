package models

type Datastore interface {
	GetUserProjects(*User) ([]Project, error)
}
