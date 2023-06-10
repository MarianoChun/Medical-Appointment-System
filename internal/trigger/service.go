package trigger

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"log"
)

const (
	pathToSpCreation     = "sql/triggers"
	creationErrorMessage = "Ocurrió un error al crear los triggers"
	creationMessage      = "Creando triggers"
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

	err := kit.ExecuteFunctionsCreation(pathToSpCreation, s.db.App())
	if err != nil {
		log.Fatalln(creationErrorMessage, err)
		return err
	}

	return nil
}
