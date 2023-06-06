package pk

import (
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	pathToPkCreation     = "sql/pk/create.sql"
	creationErrorMessage = "Ocurrió un error al crear las primary keys"
	creationMessage      = "Creando Primary Keys para las relaciones"
	pathToPkDeletion     = "sql/pk/remove.sql"
	deletionErrorMessage = "Ocurrió un error al eliminar las primary keys"
	deletionMessage      = "Eliminando Primary Keys para las relaciones"
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

	s.Delete()

	err := kit.ExecuteScript(pathToPkCreation, s.db.App())
	if err != nil {
		log.Fatalln(creationErrorMessage, err)
		return
	}
}

func (s Service) Delete() {
	log.Println(deletionMessage)

	err := kit.ExecuteScript(pathToPkDeletion, s.db.App())
	if err != nil {
		log.Fatalln(deletionErrorMessage, err)
		return
	}
}
