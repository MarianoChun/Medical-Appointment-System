package kit

import (
	"fmt"
	"log"
)

func ShowStoredProcedures(db Database) {
	title := ` 	
                                      _                                          _                            
         _                           | |                                        | |                           
  ___  _| |_  ___    ____  _____   __| |   ____    ____  ___    ____  _____   __| | _   _   ____  _____   ___ 
 /___)(_   _)/ _ \  / ___)| ___ | / _  |  |  _ \  / ___)/ _ \  / ___)| ___ | / _  || | | | / ___)| ___ | /___)
|___ |  | |_| |_| || |    | ____|( (_| |  | |_| || |   | |_| |( (___ | ____|( (_| || |_| || |    | ____||___ |
(___/    \__)\___/ |_|    |_____) \____|  |  __/ |_|    \___/  \____)|_____) \____||____/ |_|    |_____)(___/ 
                                          |_|                                                                 

	`
	options := []string{title, "1. Generar turnos disponibles", "2. Atender turnos"}
	executing := true

	for executing {
		printOptions(options)
		option, err := scanOptionSelected("Seleccione una opcion")
		if err != nil {
			log.Fatalln(err)
			return
		}

		switch option {
		case "1":
			AppointmentDateGenerator(db)
			break
		case "2":
			AppointmentAttender(db)
			break
		default:
			executing = false
			break
		}
	}
}

func printOptions(options []string) {
	for _, v := range options {
		println(v)
	}
	println("Cualquier tecla para salir")
}

func scanOptionSelected(msg string) (string, error) {
	println(msg)
	var optionSelected string
	_, err := fmt.Scanln(&optionSelected)

	if err != nil {
		fmt.Println(errorInputMessage)
		return "", err
	}

	return optionSelected, nil
}
