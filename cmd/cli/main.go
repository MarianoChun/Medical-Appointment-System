package main

import (
	"database/sql"
	"fmt"
	"github.com/boltdb/bolt"
	_ "github.com/lib/pq"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/fk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/instance"
	_ "gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/instance"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/pk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/sync"
	"log"
)

const (
	sqlDriverName     = "postgres"
	sqlDataSourceName = "user=postgres host=localhost dbname=consultorios sslmode=disable"

	boltPath = "consultorios.db"
	boltMode = 0600

	firstOption       = "1. Instanciar Base de datos"
	secondOption      = "2. Crear PK's"
	thirtyOption      = "3. Eliminar PK's"
	quarterOption     = "4. Crear FK's"
	fifthOption       = "5. Eliminar FK's"
	sixthOption       = "6. Sincronizar NoSQL con SQL"
	exitOption        = "Para salir presione cualquier tecla"
	errorInputMessage = "Ocurrió un error, intente nuevamente"
	optionMessage     = "Ingrese una opcion:"
	inputMessage      = "%s"
)

func main() {
	postgresConnection := createPostgresDatabaseConnection()
	boltConnection := createBoltDatabaseConnection()

	databaseInstantiator := instance.NewDatabaseInstantiator(postgresConnection)
	primaryKeysCreator := pk.NewPrimaryKeysCreator(postgresConnection)
	primaryKeysDeleter := pk.NewPrimaryKeysDeleter(postgresConnection)
	foreignKeysCreator := fk.NewForeignKeysCreator(postgresConnection)
	foreignKeysDeleter := fk.NewForeignKeysDeleter(postgresConnection)
	databasesSynchronizer := sync.NewDatabasesSynchronizer(postgresConnection, boltConnection)

	for {
		printOptions()

		optionSelected, err := scanOptionSelected()
		if err != nil {
			break
		}

		continueExecution := executeUseCases(optionSelected,
			databaseInstantiator,
			primaryKeysCreator,
			primaryKeysDeleter,
			foreignKeysCreator,
			foreignKeysDeleter,
			databasesSynchronizer)

		if !continueExecution {
			break
		}
	}
}

func createBoltDatabaseConnection() *bolt.DB {
	db, err := bolt.Open(boltPath, boltMode, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}

func createPostgresDatabaseConnection() *sql.DB {
	db, err := sql.Open(sqlDriverName, sqlDataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}

func printOptions() {
	fmt.Println(firstOption)
	fmt.Println(secondOption)
	fmt.Println(thirtyOption)
	fmt.Println(quarterOption)
	fmt.Println(fifthOption)
	fmt.Println(sixthOption)
	fmt.Println(exitOption)
	fmt.Println(optionMessage)
}

func executeUseCases(optionSelected string,
	databaseInstantiator instance.DatabaseInstantiator,
	primaryKeysCreator pk.PrimaryKeysCreator,
	primaryKeysDeleter pk.PrimaryKeysDeleter,
	foreignKeysCreator fk.ForeignKeysCreator,
	foreignKeysDeleter fk.ForeignKeysDeleter,
	databasesSynchronizer sync.DatabasesSynchronizer) bool {
	switch optionSelected {
	case "1":
		databaseInstantiator.Execute()
		return true
	case "2":
		primaryKeysCreator.Execute()
		return true
	case "3":
		primaryKeysDeleter.Execute()
		return true
	case "4":
		foreignKeysCreator.Execute()
		return true
	case "5":
		foreignKeysDeleter.Execute()
		return true
	case "6":
		databasesSynchronizer.Execute()
		return true
	default:
		return false
	}
}

func scanOptionSelected() (string, error) {
	var optionSelected string
	_, err := fmt.Scanf(inputMessage, &optionSelected)

	if err != nil {
		fmt.Println(errorInputMessage)
		return "", err
	}

	return optionSelected, nil
}
