package pk

import (
	"fmt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

type PrimaryKeysCreator struct {
	db kit.Database
}

func NewPrimaryKeysCreator(db kit.Database) PrimaryKeysCreator {
	return PrimaryKeysCreator{
		db: db,
	}
}

func (s PrimaryKeysCreator) Execute() {
	fmt.Println("Creating PK's")
}
