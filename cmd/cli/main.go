package main

import (
	"fmt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/cmd/cli/app"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"log"
	"strconv"
)

const (
	welcomeMessage    = "  ____                      _ _             _           \n / ___|___  _ __  ___ _   _| | |_ ___  _ __(_) ___  ___ \n| |   / _ \\| '_ \\/ __| | | | | __/ _ \\| '__| |/ _ \\/ __|\n| |__| (_) | | | \\__ \\ |_| | | || (_) | |  | | (_) \\__ \\\n \\____\\___/|_| |_|___/\\__,_|_|\\__\\___/|_|  |_|\\___/|___/\n                                                        \n    _    ____  __  __ ___ _   _ \n   / \\  |  _ \\|  \\/  |_ _| \\ | |\n  / _ \\ | | | | |\\/| || ||  \\| |\n / ___ \\| |_| | |  | || || |\\  |\n/_/   \\_\\____/|_|  |_|___|_| \\_|\n                                \n"
	spMessage         = "     _                     _ \n ___| |_ ___  _ __ ___  __| |\n/ __| __/ _ \\| '__/ _ \\/ _` |\n\\__ \\ || (_) | | |  __/ (_| |\n|___/\\__\\___/|_|  \\___|\\__,_|\n                             \n                              _                     \n _ __  _ __ ___   ___ ___  __| |_   _ _ __ ___  ___ \n| '_ \\| '__/ _ \\ / __/ _ \\/ _` | | | | '__/ _ \\/ __|\n| |_) | | | (_) | (_|  __/ (_| | |_| | | |  __/\\__ \\\n| .__/|_|  \\___/ \\___\\___|\\__,_|\\__,_|_|  \\___||___/\n|_|                                                 \n"
	mainFirstOption   = "1. Instanciar Base de datos"
	mainSecondOption  = "2. Crear PK's"
	mainThirtyOption  = "3. Eliminar PK's"
	mainQuarterOption = "4. Crear FK's"
	mainFifthOption   = "5. Eliminar FK's"
	mainSixthOption   = "6. Sincronizar NoSQL con SQL"
	mainSeventhOption = "7. Ejecutar Stored Procedures"

	spFirstOption   = "1. Generar turnos disponibles"
	spSecondOption  = "2. Atender turnos"
	spThirtyOption  = "3. Cancelar turnos"
	spQuarterOption = "4. Reservar turnos"
	spFifthOption   = "5. Generar liquidación para obras sociales"
)

func main() {
	newApp, err := app.NewApp()
	if err != nil {
		return
	}

	executeMainOptions(newApp)
}

func executeMainOptions(app app.App) {
	for {
		kit.PrintOptions(welcomeMessage, mainFirstOption, mainSecondOption, mainThirtyOption, mainQuarterOption, mainFifthOption, mainSixthOption, mainSeventhOption)

		optionSelected, err := kit.ScanOptionSelected()
		if err != nil {
			app.Close()
			break
		}

		continueExecution := executeUseCases(optionSelected, app)

		if !continueExecution {
			app.Close()
			break
		}
	}
}

func executeUseCases(optionSelected string, app app.App) bool {
	switch optionSelected {
	case "1":
		app.DatabaseService.Init()
		return true
	case "2":
		app.PrimaryKeysService.Create()
		return true
	case "3":
		app.PrimaryKeysService.Delete()
		return true
	case "4":
		app.ForeignKeysService.Create()
		return true
	case "5":
		app.PrimaryKeysService.Delete()
		return true
	case "6":
		app.DatabaseService.SyncBetweenSQLAndNoSQL()
		return true
	case "7":
		showStoredProcedures(app)
		return true
	default:
		return false
	}
}

func showStoredProcedures(app app.App) {
	executing := true

	for executing {
		kit.PrintOptions(spMessage, spFirstOption, spSecondOption, spThirtyOption, spQuarterOption, spFifthOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return
		}

		executing = executeStoredProcedures(option, app)
	}
}

func executeStoredProcedures(optionSelected string, app app.App) bool {
	switch optionSelected {
	case "1":
		executeAppointmentGenerator(app)
		return true
	case "2":
		executeAppointmentAttender(app)
		return true
	case "3":
		executeAppointmentCanceller(app)
		return true
	case "4":
		executeAppointmentReserver(app)
		return true
	case "5":
		executeInsuranceSettlementGenerator(app)
		return true
	default:
		return false
	}
}

func executeAppointmentGenerator(app app.App) {
	date, err := kit.ScanMonthAndYear()
	if err != nil {
		log.Fatalln(err)
		return
	}

	app.Appointment.Generate(date.Year(), date.Month())
}

func executeAppointmentAttender(app app.App) {
	appointmentNumberStr, err := kit.ScanOptionSelectedWithMessage("Indique el nro de turno")
	if err != nil {
		log.Fatalln(err)
		return
	}

	appointmentNumber, err := strconv.Atoi(appointmentNumberStr)
	if err != nil {
		log.Fatalln(err)
		return
	}

	app.Appointment.Attend(appointmentNumber)
}

func executeInsuranceSettlementGenerator(app app.App) {
	app.InsuranceService.GenerateSettlements()
}

func executeAppointmentReserver(app app.App) {
	clinicHistoryNumberStr, err := kit.ScanOptionSelectedWithMessage("Indique el nro de historia clinica")
	if err != nil {
		log.Fatalln(err)
		return
	}

	clinicHistoryNumber, err := strconv.Atoi(clinicHistoryNumberStr)
	if err != nil {
		log.Fatalln(err)
		return
	}

	dniStr, err := kit.ScanOptionSelectedWithMessage("Indique el dni")
	if err != nil {
		log.Fatalln(err)
		return
	}

	dni, err := strconv.Atoi(dniStr)
	if err != nil {
		log.Fatalln(err)
		return
	}

	date, err := kit.ScanDate()
	if err != nil {
		log.Fatalln(err)
		return
	}

	app.Appointment.Reserve(clinicHistoryNumber, dni, date)
}

func executeAppointmentCanceller(app app.App) {
	fmt.Println("A continuacion indicará la fecha desde")
	dateFrom, err := kit.ScanDate()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("A continuacion indicará la fecha hasta")
	dateTo, err := kit.ScanDate()
	if err != nil {
		log.Fatalln(err)
		return
	}

	dniStr, err := kit.ScanOptionSelectedWithMessage("Indique el dni")
	if err != nil {
		log.Fatalln(err)
		return
	}

	dni, err := strconv.Atoi(dniStr)
	if err != nil {
		log.Fatalln(err)
		return
	}

	app.Appointment.Cancel(dni, dateFrom, dateTo)
}
