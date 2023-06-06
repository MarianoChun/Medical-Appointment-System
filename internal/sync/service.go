package sync

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

type DatabasesSynchronizer struct {
	db kit.Database
}

func NewDatabasesSynchronizer(db kit.Database) DatabasesSynchronizer {
	return DatabasesSynchronizer{
		db: db,
	}
}

func (s DatabasesSynchronizer) Execute() {
}
