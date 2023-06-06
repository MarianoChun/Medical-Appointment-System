package fk

import (
	"fmt"
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	pathToFkDeletion       = "sql/fk/remove.sql"
	deletionMessage        = "Eliminando foreign keys para las relaciones"
	deletionSuccessMessage = "Foreign keys eliminadas correctamente!"
	deletionErrorMessage   = "Ocurrió un error al eliminar las foreign keys"
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
	fmt.Println(deletionMessage)

	err := kit.ExecuteScript(pathToFkDeletion, s.db.App())

	if err != nil {
		log.Fatalln(deletionErrorMessage, err)
		return
	}

	fmt.Println(deletionSuccessMessage)
}
