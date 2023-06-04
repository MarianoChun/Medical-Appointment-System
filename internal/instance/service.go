package instance

import (
	"database/sql"
	"fmt"
)

type DatabaseInstantiator struct {
	db *sql.DB
}

func NewDatabaseInstantiator(db *sql.DB) DatabaseInstantiator {
	return DatabaseInstantiator{
		db: db,
	}
}

func (s DatabaseInstantiator) Execute() {
	fmt.Println("Instantiating Database")
}
