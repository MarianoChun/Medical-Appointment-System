package fk

import (
	"fmt"
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	pathToFkCreation       = "sql/fk/create.sql"
	creationMessage        = "Creando foreign keys para las relaciones"
	creationSuccessMessage = "Foreign keys creadas correctamente!"
	creationErrorMessage   = "Ocurrió un error al crear las foreign keys"
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
	fmt.Println(creationMessage)

	err := kit.ExecuteScript(pathToFkCreation, s.db.App())

	if err != nil {
		log.Fatalln(creationErrorMessage, err)
		return
	}

	fmt.Println(creationSuccessMessage)
}
