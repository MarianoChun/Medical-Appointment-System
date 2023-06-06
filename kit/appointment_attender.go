package kit

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func AppointmentAttender(db Database) {
	options := []string{"1. Atender turno"}
	var executing = true

	for executing {
		PrintOptions("", options...)

		option, err := ScanOptionSelectedWithMessage("Seleccione una opcion")
		if err != nil {
			log.Fatalln(err)
			return
		}

		switch option {
		case "1":
			attendAppointment(db.App())
			break
		default:
			executing = false
			break
		}

	}

}

func attendAppointment(db *sql.DB) {
	app, err := ScanOptionSelectedWithMessage("Ingrese el numero de turno a atender")
	if err != nil {
		db.Close()
	}

	intApp, err := strconv.Atoi(app)
	if err != nil {
		fmt.Println("Ocurrio un error al procesar el numero de turno")
		return
	}

	var attentionResult bool
	db.QueryRow(`select attend_appointment($1);`, intApp).Scan(&attentionResult)
	if err != nil {
		log.Fatal(err)
	}

	if attentionResult {
		fmt.Printf("Turno %v atendido con exito. \n", intApp)
	} else {
		fmt.Printf("Ocurrió un error al atender el turno nro %v. \n", intApp)
	}

}
