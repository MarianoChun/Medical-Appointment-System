package pk

import (
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	pathToPkDeletion     = "sql/pk/remove.sql"
	deletionErrorMessage = "Ocurrió un error al eliminar las primary keys"
	deletionMessage      = "Eliminando Primary Keys para las relaciones"
)

type Deleter struct {
	db kit.Database
}

func NewPrimaryKeysDeleter(db kit.Database) Deleter {
	return Deleter{
		db: db,
	}
}

func (s Deleter) Execute() {
	log.Println(deletionMessage)

	err := kit.ExecuteScript(pathToPkDeletion, s.db.App())
	if err != nil {
		log.Fatalln(deletionErrorMessage, err)
		return
	}
}
