package pk

import (
	"fmt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

type PrimaryKeysDeleter struct {
	db kit.Database
}

func NewPrimaryKeysDeleter(db kit.Database) PrimaryKeysDeleter {
	return PrimaryKeysDeleter{
		db: db,
	}
}

func (s PrimaryKeysDeleter) Execute() {
	fmt.Println("Deleting PK's")
}
