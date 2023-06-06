package fk

import (
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	pathToFkDeletion       = "sql/fk/remove.sql"
	deletionMessage        = "Eliminando foreign keys para las relaciones"
	deletionSuccessMessage = "Foreign keys eliminadas correctamente!"
	deletionErrorMessage   = "Ocurrió un error al eliminar las foreign keys"
)

type Deleter struct {
	db kit.Database
}

func NewForeignKeysDeleter(db kit.Database) Deleter {
	return Deleter{
		db: db,
	}
}

func (s Deleter) Execute() {
	log.Println(deletionMessage)

	err := kit.ExecuteScript(pathToFkDeletion, s.db.App())

	if err != nil {
		log.Fatalln(deletionErrorMessage, err)
		return
	}

	log.Println(deletionSuccessMessage)
}
