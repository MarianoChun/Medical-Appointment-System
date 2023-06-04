package fk

import (
	"database/sql"
	"fmt"
)

type ForeignKeysDeleter struct {
	db *sql.DB
}

func NewForeignKeysDeleter(db *sql.DB) ForeignKeysDeleter {
	return ForeignKeysDeleter{
		db: db,
	}
}

func (s ForeignKeysDeleter) Execute() {
	fmt.Println("Deleting FK's")
}
