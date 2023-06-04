package pk

import (
	"database/sql"
	"fmt"
)

type PrimaryKeysCreator struct {
	db *sql.DB
}

func NewPrimaryKeysCreator(db *sql.DB) PrimaryKeysCreator {
	return PrimaryKeysCreator{
		db: db,
	}
}

func (s PrimaryKeysCreator) Execute() {
	fmt.Println("Creating PK's")
}
