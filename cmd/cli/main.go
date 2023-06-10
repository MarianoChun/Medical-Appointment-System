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
	mainSeventhOption = "7. Ver datos NoSQL"
	mainEightOption   = "8. Ejecutar Stored Procedures"

	spFirstOption   = "1. Generar turnos disponibles"
	spSecondOption  = "2. Atender turnos"
	spThirtyOption  = "3. Cancelar turnos"
	spQuarterOption = "4. Reservar turnos"
	spFifthOption   = "5. Generar liquidación para obras sociales"
	spSixthOption   = "6. Enviar emails de inasistencia"
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
		kit.PrintOptions(welcomeMessage, mainFirstOption, mainSecondOption, mainThirtyOption, mainQuarterOption, mainFifthOption, mainSixthOption, mainSeventhOption, mainEightOption)

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
		return app.DatabaseService.Init() == nil
	case "2":
		return app.PrimaryKeysService.Create() == nil
	case "3":
		return app.PrimaryKeysService.Delete() == nil
	case "4":
		return app.ForeignKeysService.Create() == nil
	case "5":
		return app.PrimaryKeysService.Delete() == nil
	case "6":
		return app.DatabaseService.SyncBetweenSQLAndNoSQL() == nil
	case "7":
		return app.DatabaseService.ViewNoSQL() == nil
	case "8":
		return showStoredProcedures(app) == nil
	default:
		return false
	}
}

func showStoredProcedures(app app.App) error {
	executing := true

	for executing {
		kit.PrintOptions(spMessage, spFirstOption, spSecondOption, spThirtyOption, spQuarterOption, spFifthOption, spSixthOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		executing = executeStoredProcedures(option, app)
	}

	return nil
}

func executeStoredProcedures(optionSelected string, app app.App) bool {
	switch optionSelected {
	case "1":
		return executeAppointmentGenerator(app) == nil
	case "2":
		return executeAppointmentAttender(app) == nil
	case "3":
		return executeAppointmentCanceller(app) == nil
	case "4":
		return executeAppointmentReserver(app) == nil
	case "5":
		return app.InsuranceService.GenerateSettlements() == nil
	case "6":
		return app.EmailService.SendAbsenseEmails() == nil
	default:
		return false
	}
}

func executeAppointmentGenerator(app app.App) error {
	date, err := kit.ScanMonthAndYear()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return app.Appointment.Generate(date.Year(), int(date.Month()))
}

func executeAppointmentAttender(app app.App) error {
	appointmentNumberStr, err := kit.ScanOptionSelectedWithMessage("Indique el nro de turno")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	appointmentNumber, err := strconv.Atoi(appointmentNumberStr)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return app.Appointment.Attend(appointmentNumber)
}

func executeAppointmentReserver(app app.App) error {
	clinicHistoryNumberStr, err := kit.ScanOptionSelectedWithMessage("Indique el nro de historia clinica del paciente")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	clinicHistoryNumber, err := strconv.Atoi(clinicHistoryNumberStr)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	dniStr, err := kit.ScanOptionSelectedWithMessage("Indique el dni del medique")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	dni, err := strconv.Atoi(dniStr)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	date, err := kit.ScanDateAndHour()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return app.Appointment.Reserve(clinicHistoryNumber, dni, date)
}

func executeAppointmentCanceller(app app.App) error {
	fmt.Println("A continuacion indicará la fecha desde")
	dateFrom, err := kit.ScanDate()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	fmt.Println("A continuacion indicará la fecha hasta")
	dateTo, err := kit.ScanDate()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	dniStr, err := kit.ScanOptionSelectedWithMessage("Indique el dni")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	dni, err := strconv.Atoi(dniStr)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return app.Appointment.Cancel(dni, dateFrom, dateTo)
}
