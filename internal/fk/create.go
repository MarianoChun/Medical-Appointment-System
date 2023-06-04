package fk

import (
	"database/sql"
	"fmt"
)

type ForeignKeysCreator struct {
	db *sql.DB
}

func NewForeignKeysCreator(db *sql.DB) ForeignKeysCreator {
	return ForeignKeysCreator{
		db: db,
	}
}

func (s ForeignKeysCreator) Execute() {
	fmt.Println("Creating FK's")
}
