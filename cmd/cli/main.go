package main

import (
	"fmt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/fk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/instance"
	_ "gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/instance"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/pk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/sync"
)

const (
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
	databaseInstantiator := instance.NewDatabaseInstantiator()
	primaryKeysService := pk.NewPrimaryKeysService()
	foreignKeysService := fk.NewForeignKeysService()
	databasesSynchronizer := sync.NewDatabasesSynchronizer()

	for {
		printOptions()

		optionSelected, err := scanOptionSelected()
		if err != nil {
			break
		}

		continueExecution := executeUseCases(optionSelected, databaseInstantiator, primaryKeysService, foreignKeysService, databasesSynchronizer)

		if !continueExecution {
			break
		}
	}
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

func executeUseCases(optionSelected string, instantiator instance.DatabaseInstantiatorService, primaryKeysService pk.PrimaryKeysService, foreignKeysService fk.ForeignKeysService, synchronizer sync.DatabasesSynchronizerService) bool {
	switch optionSelected {
	case "1":
		instantiator.Execute()
		return true
	case "2":
		primaryKeysService.Create()
		return true
	case "3":
		primaryKeysService.Delete()
		return true
	case "4":
		foreignKeysService.Create()
		return true
	case "5":
		foreignKeysService.Delete()
		return true
	case "6":
		synchronizer.Execute()
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
