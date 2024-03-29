package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/cmd/cli/app"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

const (
	welcomeMessage = "  ____                      _ _             _           \n / ___|___  _ __  ___ _   _| | |_ ___  _ __(_) ___  ___ \n| |   / _ \\| '_ \\/ __| | | | | __/ _ \\| '__| |/ _ \\/ __|\n| |__| (_) | | | \\__ \\ |_| | | || (_) | |  | | (_) \\__ \\\n \\____\\___/|_| |_|___/\\__,_|_|\\__\\___/|_|  |_|\\___/|___/\n                                                        \n    _    ____  __  __ ___ _   _ \n   / \\  |  _ \\|  \\/  |_ _| \\ | |\n  / _ \\ | | | | |\\/| || ||  \\| |\n / ___ \\| |_| | |  | || || |\\  |\n/_/   \\_\\____/|_|  |_|___|_| \\_|\n                                \n"
	dbMessage      = " ____   ___  _     \n/ ___| / _ \\| |    \n\\___ \\| | | | |    \n ___) | |_| | |___ \n|____/ \\__\\_\\_____|\n"
	spMessage      = "     _                     _ \n ___| |_ ___  _ __ ___  __| |\n/ __| __/ _ \\| '__/ _ \\/ _` |\n\\__ \\ || (_) | | |  __/ (_| |\n|___/\\__\\___/|_|  \\___|\\__,_|\n                             \n                              _                     \n _ __  _ __ ___   ___ ___  __| |_   _ _ __ ___  ___ \n| '_ \\| '__/ _ \\ / __/ _ \\/ _` | | | | '__/ _ \\/ __|\n| |_) | | | (_) | (_|  __/ (_| | |_| | | |  __/\\__ \\\n| .__/|_|  \\___/ \\___\\___|\\__,_|\\__,_|_|  \\___||___/\n|_|                                                 \n"
	noSqlMessage   = " _   _      ____   ___  _     \n| \\ | | ___/ ___| / _ \\| |    \n|  \\| |/ _ \\___ \\| | | | |    \n| |\\  | (_) |__) | |_| | |___ \n|_| \\_|\\___/____/ \\__\\_\\_____|\n                              \n"

	mainFirstOption  = "1.  Administración Base de datos SQL"
	mainSecondOption = "2.  Administración de NoSQL"

	dbFirstOption   = "1. Crear Base de datos"
	dbSecondOption  = "2. Crear Tablas"
	dbThirdOption   = "3. Crear Primary/Foreign Keys"
	dbQuarterOption = "4. Eliminar Primary/Foreign Keys"
	dbFifthOption   = "5. Crear de Stored Procedures"
	dbSixthOption   = "6. Crear Triggers"
	dbSeventhOption = "7. Insertar data"
	dbEighthOption  = "8. Ejecutar Stored Procedures"

	noSqlFirstOption  = "1. Sincronizar Bases de datos"
	noSqlSecondOption = "2. Ver datos"

	spFirstOption   = "1. Generar turnos disponibles"
	spSecondOption  = "2. Atender turnos"
	spThirtyOption  = "3. Cancelar turnos"
	spQuarterOption = "4. Reservar turnos"
	spFifthOption   = "5. Generar liquidación para obras sociales"
	spSixthOption   = "6. Enviar emails de inasistencia"
	spSevenOption   = "7. Enviar emails de recordatorio"
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
		kit.PrintOptions(welcomeMessage, mainFirstOption, mainSecondOption)

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
		return showSQL(app) == nil
	case "2":
		return showNoSql(app) == nil
	default:
		return false
	}
}

func showSQL(app app.App) error {
	executing := true

	for executing {
		kit.PrintOptions(dbMessage, dbFirstOption, dbSecondOption, dbThirdOption, dbQuarterOption, dbFifthOption, dbSixthOption, dbSeventhOption, dbEighthOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		switch option {
		case "1":
			executing = app.DatabaseService.Create() == nil
		case "2":
			executing = app.DatabaseService.CreateTables() == nil
		case "3":
			executing = app.PrimaryKeysService.Create() == nil
			executing = app.ForeignKeysService.Create() == nil
		case "4":
			executing = app.PrimaryKeysService.Delete() == nil
			executing = app.ForeignKeysService.Delete() == nil
		case "5":
			executing = app.StoredProcedureService.Create() == nil
		case "6":
			executing = app.TriggerService.Create() == nil
		case "7":
			executing = app.DatabaseService.InsertData() == nil
		case "8":
			executing = showStoredProcedures(app) == nil
		default:
			executing = false
		}
	}

	return nil
}

func showStoredProcedures(app app.App) error {
	executing := true

	for executing {
		kit.PrintOptions(spMessage, spFirstOption, spSecondOption, spThirtyOption, spQuarterOption, spFifthOption, spSixthOption, spSevenOption)
		option, err := kit.ScanOptionSelected()
		if err != nil {
			log.Fatalln(err)
			return err
		}

		switch option {
		case "1":
			executing = executeAppointmentGenerator(app) == nil
		case "2":
			executing = executeAppointmentAttender(app) == nil
		case "3":
			executing = executeAppointmentCanceller(app) == nil
		case "4":
			executing = executeAppointmentReserver(app) == nil
		case "5":
			executing = app.InsuranceService.GenerateSettlements() == nil
		case "6":
			executing = app.EmailService.SendAbsenseEmails() == nil
		case "7":
			executing = app.EmailService.SendReminderEmails() == nil
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
	currentMonth := time.Now().Format("1")
	currentYear := time.Now().Format("2006")

	_, err := app.GetDb().App().Exec("delete from turno;")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	query := fmt.Sprintf("select generate_appointments_in_month(%s, %s);", currentYear, currentMonth)
	_, err = app.GetDb().App().Exec(query)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	appointmentRequests, err := kit.QueryRowsFromTable("solicitud_reservas", "nro_orden", app.GetDb().App())
	defer appointmentRequests.Close()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	for appointmentRequests.Next() {

		var appointment internal.Appointment
		var unusedColumn int
		var reserveDate time.Time
		var reserveHour time.Time

		if err := appointmentRequests.Scan(&unusedColumn, &appointment.PatientNumber, &appointment.MedicDni, &reserveDate, &reserveHour); err != nil {
			log.Fatalln(err)
			return err
		}

		appointmentTimestamp := time.Date(reserveDate.Year(), reserveDate.Month(), reserveDate.Day(), reserveHour.Hour(), reserveHour.Minute(), 0, 0, time.UTC)
		err := app.Appointment.Reserve(appointment.PatientNumber, appointment.MedicDni, appointmentTimestamp)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	}
	return nil
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
