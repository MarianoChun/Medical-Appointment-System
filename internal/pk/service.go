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

func (s Service) Create() error {
	log.Println(creationMessage)

	err := s.Delete()
	if err != nil {
		return err
	}

	err = kit.ExecuteScript(pathToPkCreation, s.db.App())
	if err != nil {
		log.Fatalln(creationErrorMessage, err)
		return err
	}

	return nil
}

func (s Service) Delete() error {
	log.Println(deletionMessage)

	err := kit.ExecuteScript(pathToPkDeletion, s.db.App())
	if err != nil {
		log.Fatalln(deletionErrorMessage, err)
		return err
	}

	return nil
}
