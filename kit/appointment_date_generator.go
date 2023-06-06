package kit

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func AppointmentDateGenerator(db Database) {
	options := []string{"1. Crear turnos disponibles"}
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
			generateAppointments(db.App())
			break
		default:
			executing = false
			break
		}

	}

}

func generateAppointments(db *sql.DB) {
	year, err := ScanOptionSelectedWithMessage("Seleccione un año para generar")
	if err != nil {
		db.Close()
	}

	intYear, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("Ocurrio un error al procesar el año")
		return
	}

	month, err := ScanOptionSelectedWithMessage("Seleccione un numero de mes para generar, por ejemplo para junio seleccione 6")
	if err != nil {
		db.Close()
	}

	intMonth, err := strconv.Atoi(month)
	if err != nil {
		fmt.Println("Ocurrio un error al procesar el mes")
		return
	}

	var generationResult bool
	db.QueryRow(`select generate_appointments_in_month($1, $2);`, intYear, intMonth).Scan(&generationResult)
	if err != nil {
		log.Fatal(err)
	}

	if generationResult {
		fmt.Printf("Generando turnos para el %v/%v. \n", month, year)
	} else {
		fmt.Printf("Los turnos diponibles para el %v/%v ya se encuentran generados. \n", year, month)
	}

}
