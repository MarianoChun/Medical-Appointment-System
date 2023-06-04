package pk

import (
	"database/sql"
	"fmt"
)

type PrimaryKeysDeleter struct {
	db *sql.DB
}

func NewPrimaryKeysDeleter(db *sql.DB) PrimaryKeysDeleter {
	return PrimaryKeysDeleter{
		db: db,
	}
}

func (s PrimaryKeysDeleter) Execute() {
	fmt.Println("Deleting PK's")
}
