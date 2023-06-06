package pk

import (
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	pathToPkCreation     = "sql/pk/create.sql"
	creationErrorMessage = "Ocurrió un error al crear las primary keys"
	creationMessage      = "Creando Primary Keys para las relaciones"
)

type Creator struct {
	db kit.Database
}

func NewPrimaryKeysCreator(db kit.Database) Creator {
	return Creator{
		db: db,
	}
}

func (s Creator) Execute() {
	log.Println(creationMessage)

	err := kit.ExecuteScript(pathToPkCreation, s.db.App())
	if err != nil {
		log.Fatalln(creationErrorMessage, err)
		return
	}
}
