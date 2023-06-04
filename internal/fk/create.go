package fk

import (
	"fmt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

type ForeignKeysCreator struct {
	db kit.Database
}

func NewForeignKeysCreator(db kit.Database) ForeignKeysCreator {
	return ForeignKeysCreator{
		db: db,
	}
}

func (s ForeignKeysCreator) Execute() {
	fmt.Println("Creating FK's")
}
