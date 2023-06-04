package sync

import (
	"database/sql"
	"fmt"
	"github.com/boltdb/bolt"
)

type DatabasesSynchronizer struct {
	postgres *sql.DB
	bolt     *bolt.DB
}

func NewDatabasesSynchronizer(postgres *sql.DB, bolt *bolt.DB) DatabasesSynchronizer {
	return DatabasesSynchronizer{
		postgres: postgres,
		bolt:     bolt,
	}
}

func (s DatabasesSynchronizer) Execute() {
	fmt.Println("Sync Database")
}
