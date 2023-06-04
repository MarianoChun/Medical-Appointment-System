package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/fk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/instance"
	_ "gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/instance"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/pk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/sync"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"time"
)

const (
	welcomeMessage    = "  ____                      _ _             _           \n / ___|___  _ __  ___ _   _| | |_ ___  _ __(_) ___  ___ \n| |   / _ \\| '_ \\/ __| | | | | __/ _ \\| '__| |/ _ \\/ __|\n| |__| (_) | | | \\__ \\ |_| | | || (_) | |  | | (_) \\__ \\\n \\____\\___/|_| |_|___/\\__,_|_|\\__\\___/|_|  |_|\\___/|___/\n                                                        \n    _    ____  __  __ ___ _   _ \n   / \\  |  _ \\|  \\/  |_ _| \\ | |\n  / _ \\ | | | | |\\/| || ||  \\| |\n / ___ \\| |_| | |  | || || |\\  |\n/_/   \\_\\____/|_|  |_|___|_| \\_|\n                                \n"
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
	database, err := kit.NewDatabase()
	if err != nil {
		return
	}

	initializer := instance.NewDatabaseInitializer(database)
	primaryKeysCreator := pk.NewPrimaryKeysCreator(database)
	primaryKeysDeleter := pk.NewPrimaryKeysDeleter(database)
	foreignKeysCreator := fk.NewForeignKeysCreator(database)
	foreignKeysDeleter := fk.NewForeignKeysDeleter(database)
	databasesSynchronizer := sync.NewDatabasesSynchronizer(database)

	for {
		printOptions()

		optionSelected, err := scanOptionSelected()
		if err != nil {
			database.Close()
			break
		}

		continueExecution := executeUseCases(optionSelected,
			initializer,
			primaryKeysCreator,
			primaryKeysDeleter,
			foreignKeysCreator,
			foreignKeysDeleter,
			databasesSynchronizer)

		if !continueExecution {
			database.Close()
			break
		}
	}
}

func printOptions() {
	fmt.Println(welcomeMessage)
	time.Sleep(1 * time.Second)
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
	initializer instance.DatabaseInitializer,
	primaryKeysCreator pk.PrimaryKeysCreator,
	primaryKeysDeleter pk.PrimaryKeysDeleter,
	foreignKeysCreator fk.ForeignKeysCreator,
	foreignKeysDeleter fk.ForeignKeysDeleter,
	databasesSynchronizer sync.DatabasesSynchronizer) bool {
	switch optionSelected {
	case "1":
		initializer.Execute()
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
