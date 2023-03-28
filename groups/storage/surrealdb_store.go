package groups

import (
	sdb "github.com/iKaio/surrealdb.go"
	"groups/groups_errors.go"
)

// SurrealDB store details
type SurrealDB_store struct {
	Environment string
	location string
	db sdb.DB
	ns string
	tbl string
}

func (*SurrealDB_store) FindByName(name string) (*Group, error) {
	g *Group := NewGroup()
	return g, nil
}