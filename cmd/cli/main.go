package main

import (
	"fmt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/cmd/cli/app"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"log"
	"strconv"
)

const (
	welcomeMessage = "  ____                      _ _             _           \n / ___|___  _ __  ___ _   _| | |_ ___  _ __(_) ___  ___ \n| |   / _ \\| '_ \\/ __| | | | | __/ _ \\| '__| |/ _ \\/ __|\n| |__| (_) | | | \\__ \\ |_| | | || (_) | |  | | (_) \\__ \\\n \\____\\___/|_| |_|___/\\__,_|_|\\__\\___/|_|  |_|\\___/|___/\n                                                        \n    _    ____  __  __ ___ _   _ \n   / \\  |  _ \\|  \\/  |_ _| \\ | |\n  / _ \\ | | | | |\\/| || ||  \\| |\n / ___ \\| |_| | |  | || || |\\  |\n/_/   \\_\\____/|_|  |_|___|_| \\_|\n                                \n"
	pkMessage      = " ____       _                              _  __              \n|  _ \\ _ __(_)_ __ ___   __ _ _ __ _   _  | |/ /___ _   _ ___ \n| |_) | '__| | '_ ` _ \\ / _` | '__| | | | | ' // _ \\ | | / __|\n|  __/| |  | | | | | | | (_| | |  | |_| | | . \\  __/ |_| \\__ \\\n|_|   |_|  |_|_| |_| |_|\\__,_|_|   \\__, | |_|\\_\\___|\\__, |___/\n                                   |___/            |___/   "
	fkMessage      = " _____              _               _  __              \n|  ___|__  _ __ ___(_) __ _ _ __   | |/ /___ _   _ ___ \n| |_ / _ \\| '__/ _ \\ |/ _` | '_ \\  | ' // _ \\ | | / __|\n|  _| (_) | | |  __/ | (_| | | | | | . \\  __/ |_| \\__ \\\n|_|  \\___/|_|  \\___|_|\\__, |_| |_| |_|\\_\\___|\\__, |___/\n                      |___/                  |___/     \n"
	tgMessage      = " _____     _                           \n|_   _| __(_) __ _  __ _  ___ _ __ ___ \n  | || '__| |/ _` |/ _` |/ _ \\ '__/ __|\n  | || |  | | (_| | (_| |  __/ |  \\__ \\\n  |_||_|  |_|\\__, |\\__, |\\___|_|  |___/\n             |___/ |___/               \n"
	spMessage      = "     _                     _ \n ___| |_ ___  _ __ ___  __| |\n/ __| __/ _ \\| '__/ _ \\/ _` |\n\\__ \\ || (_) | | |  __/ (_| |\n|___/\\__\\___/|_|  \\___|\\__,_|\n                             \n                              _                     \n _ __  _ __ ___   ___ ___  __| |_   _ _ __ ___  ___ \n| '_ \\| '__/ _ \\ / __/ _ \\/ _` | | | | '__/ _ \\/ __|\n| |_) | | | (_) | (_|  __/ (_| | |_| | | |  __/\\__ \\\n| .__/|_|  \\___/ \\___\\___|\\__,_|\\__,_|_|  \\___||___/\n|_|                                                 \n"
	noSqlMessage   = " _   _      ____   ___  _     \n| \\ | | ___/ ___| / _ \\| |    \n|  \\| |/ _ \\___ \\| | | | |    \n| |\\  | (_) |__) | |_| | |___ \n|_| \\_|\\___/____/ \\__\\_\\_____|\n                              \n"

	mainFirstOption   = "1.  Crear Base de datos"
	mainSecondOption  = "2.  Insertar datos"
	mainThirtyOption  = "3.  Administración de Primary Keys"
	mainQuarterOption = "4.  Administración de Foreign Keys"
	mainFifthOption   = "5.  Administración de Stored Procedures"
	mainSixthOption   = "6.  Administración Triggers"
	mainSeventhOption = "7.  Administración de NoSQL"

	pkFirstOption  = "1. Crear Primary Keys"
	pkSecondOption = "2. Eliminar Primary Keys"

	fkFirstOption  = "1. Crear Foreign Keys"
	fkSecondOption = "2. Eliminar Foreign Keys"

	triggerFirstOption = "1. Crear Triggers"

	noSqlFirstOption  = "1. Sincronizar Bases de datos"
	noSqlSecondOption = "2. Ver datos"

	spFirstOption   = "1. Crear Stored Procedures"
	spSecondOption  = "2. Generar turnos disponibles"
	spThirtyOption  = "3. Atender turnos"
	spQuarterOption = "4. Cancelar turnos"
	spFifthOption   = "5. Reservar turnos"
	spSixthOption   = "6. Generar liquidación para obras sociales"
	spSeventhOption = "7. Enviar emails de inasistencia"
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
		return app.DatabaseService.Create() == nil
	case "2":
		return app.DatabaseService.InsertData() == nil
	case "3":
		return showPrimaryKeys(app) == nil
	case "4":
		return showForeignKeys(app) == nil
	case "5":
		return showStoredProcedures(app) == nil
	case "6":
		return showTriggers(app) == nil
	case "7":
		return showNoSql(app) == nil
	default:
		return false
	}
}

func showPrimaryKeys(app app.App) error {
	executing := true

	for executing {
		kit.PrintOptions(pkMessage, pkFirstOption, pkSecondOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		switch option {
		case "1":
			executing = app.PrimaryKeysService.Create() == nil
		case "2":
			executing = app.PrimaryKeysService.Delete() == nil
		default:
			executing = false
		}
	}

	return nil
}

func showForeignKeys(app app.App) error {
	executing := true

	for executing {
		kit.PrintOptions(fkMessage, fkFirstOption, fkSecondOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		switch option {
		case "1":
			executing = app.ForeignKeysService.Create() == nil
		case "2":
			executing = app.ForeignKeysService.Delete() == nil
		default:
			executing = false
		}
	}

	return nil
}

func showTriggers(app app.App) error {
	executing := true

	for executing {
		kit.PrintOptions(tgMessage, triggerFirstOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		switch option {
		case "1":
			executing = app.TriggerService.Create() == nil
		default:
			executing = false
		}
	}

	return nil
}

func showStoredProcedures(app app.App) error {
	executing := true

	for executing {
		kit.PrintOptions(spMessage, spFirstOption, spSecondOption, spThirtyOption, spQuarterOption, spFifthOption, spSixthOption, spSeventhOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		switch option {
		case "1":
			executing = app.StoredProcedureService.Create() == nil
		case "2":
			executing = executeAppointmentGenerator(app) == nil
		case "3":
			executing = executeAppointmentAttender(app) == nil
		case "4":
			executing = executeAppointmentCanceller(app) == nil
		case "5":
			executing = executeAppointmentReserver(app) == nil
		case "6":
			executing = app.InsuranceService.GenerateSettlements() == nil
		case "7":
			executing = app.EmailService.SendAbsenseEmails() == nil
		default:
			executing = false
		}
	}

	return nil
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

func showNoSql(app app.App) error {
	executing := true

	for executing {
		kit.PrintOptions(noSqlMessage, noSqlFirstOption, noSqlSecondOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		switch option {
		case "1":
			executing = app.DatabaseService.SyncBetweenSQLAndNoSQL() == nil
		case "2":
			executing = app.DatabaseService.ViewNoSQL() == nil
		default:
			executing = false
		}
	}

	return nil
}
