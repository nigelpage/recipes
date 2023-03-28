package groups

import (
	uuid "github.com/google/uuid"
)

type GroupsStore interface {
	FindByName(name string) (*Group, error)
	AddOrUpdate(string, []string, uuid.UUID, bool) (*Group, error)
	Delete(string) (error)
}

type Group struct {
	Name string
	People []uuid.UUID
	Owner uuid.UUID
	Private bool
}

func NewGroup(name string, owner uuid.UUID, private bool) *Group {
	return &Group {
		Name: name,
		Owner: owner,
		Private: private,
	}
}