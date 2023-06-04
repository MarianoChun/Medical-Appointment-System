package pk

import (
	"fmt"
	"log"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	path_to_pk_deletion = "sql/pk/remove.sql"
	deletion_error_message = "Ocurrió un error al eliminar las primary keys"
	deletion_message = "Eliminando Primary Keys para las relaciones"
)

type PrimaryKeysDeleter struct {
	db kit.Database
}

func NewPrimaryKeysDeleter(db kit.Database) PrimaryKeysDeleter {
	return PrimaryKeysDeleter{
		db: db,
	}
}

func (s PrimaryKeysDeleter) Execute() {
	fmt.Println(deletion_message);
	err := kit.ExecuteScript(path_to_pk_deletion, s.db.App())
	if err != nil {
			log.Fatalln(deletion_error_message, err)
		return
	}

}
