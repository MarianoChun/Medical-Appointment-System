package pk

import (
	"fmt"
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	path_to_pk_creation = "sql/pk/create.sql"
	creation_error_message = "Ocurrió un error al crear las primary keys"
	creation_message = "Creando Primary Keys para las relaciones"
)

type PrimaryKeysCreator struct {
	db kit.Database
}

func NewPrimaryKeysCreator(db kit.Database) PrimaryKeysCreator {
	return PrimaryKeysCreator{
		db: db,
	}
}

func (s PrimaryKeysCreator) Execute() {
	fmt.Println(creation_message);
	err := kit.ExecuteScript(path_to_pk_creation, s.db.App())
	if err != nil {
		log.Fatalln(creation_error_message, err)
		return
	}

}
