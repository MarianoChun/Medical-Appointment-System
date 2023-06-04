package fk

import (
	"fmt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

type ForeignKeysDeleter struct {
	db kit.Database
}

func NewForeignKeysDeleter(db kit.Database) ForeignKeysDeleter {
	return ForeignKeysDeleter{
		db: db,
	}
}

func (s ForeignKeysDeleter) Execute() {
	fmt.Println("Deleting FK's")
}
