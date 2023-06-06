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
	pathToFkDeletion       = "sql/fk/remove.sql"
	deletionMessage        = "Eliminando foreign keys para las relaciones"
	deletionSuccessMessage = "Foreign keys eliminadas correctamente!"
	deletionErrorMessage   = "Ocurrió un error al eliminar las foreign keys"
)

type Service struct {
	db kit.Database
}

func NewService(db kit.Database) Service {
	return Service{
		db: db,
	}
}

func (s Service) Create() {
	log.Println(creationMessage)

	err := kit.ExecuteScript(pathToFkCreation, s.db.App())

	if err != nil {
		log.Fatalln(creationErrorMessage, err)
		return
	}

	log.Println(creationSuccessMessage)
}

func (s Service) Delete() {
	log.Println(deletionMessage)

	err := kit.ExecuteScript(pathToFkDeletion, s.db.App())

	if err != nil {
		log.Fatalln(deletionErrorMessage, err)
		return
	}

	log.Println(deletionSuccessMessage)
}
