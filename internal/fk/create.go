package fk

import (
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	pathToFkCreation       = "sql/fk/create.sql"
	creationMessage        = "Creando foreign keys para las relaciones"
	creationSuccessMessage = "Foreign keys creadas correctamente!"
	creationErrorMessage   = "Ocurrió un error al crear las foreign keys"
)

type Creator struct {
	db kit.Database
}

func NewForeignKeysCreator(db kit.Database) Creator {
	return Creator{
		db: db,
	}
}

func (s Creator) Execute() {
	log.Println(creationMessage)

	err := kit.ExecuteScript(pathToFkCreation, s.db.App())

	if err != nil {
		log.Fatalln(creationErrorMessage, err)
		return
	}

	log.Println(creationSuccessMessage)
}
