package appointment

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

type Service struct {
	db kit.Database
}

func NewService(db kit.Database) Service {
	return Service{
		db: db,
	}
}

func (s Service) Generate(year int, month int) error {
	begin, err := s.db.App().Begin()
	if err != nil {
		return err
	}

	// Elegimos serializable para que no se generen turnos duplicados
	_, err = begin.Exec("set transaction isolation level serializable;")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	query := fmt.Sprintf("select generate_appointments_in_month(%s, %s);", strconv.Itoa(year), strconv.Itoa(month))

	var result bool
	err = begin.QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	begin.Commit()

	if result {
		fmt.Printf("Generando turnos para el %v/%v. \n", month, year)
	} else {
		fmt.Printf("Los turnos diponibles para el %v/%v ya se encuentran generados. \n", year, month)
	}

	return nil
}

func (s Service) Attend(appointmentNumber int) error {
	// Si se ejecuta otra transaccion que modifique la tabla de turnos, la otra transaccion se bloquea
	query := fmt.Sprintf("select attend_appointment(%s);", strconv.Itoa(appointmentNumber))

	begin, err := s.db.App().Begin()
	if err != nil {
		return err
	}

	_, err = begin.Exec("set transaction isolation level repeatable read;")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	var result bool
	err = begin.QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	begin.Commit()

	if result {
		log.Printf("Turno %v atendido con exito. \n", appointmentNumber)
	} else {
		log.Printf("Ocurrió un error al atender el turno nro %v. \n", appointmentNumber)
	}

	return nil
}

func (s Service) Cancel(dni int, dateFrom time.Time, dateTo time.Time) error {
	begin, err := s.db.App().Begin()
	if err != nil {
		return err
	}

	// Si se ejecuta otra transaccion que modifique la tabla de turnos, la otra transaccion se bloquea
	_, err = begin.Exec("set transaction isolation level repeatable read;")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	query := fmt.Sprintf("select cancel_appointment(%s, '%s', '%s');", strconv.Itoa(dni), kit.TimeToDateString(dateFrom), kit.TimeToDateString(dateTo))

	var result int
	err = begin.QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	begin.Commit()

	log.Printf("Los cantidad de turnos cancelados fueron %s \n", strconv.Itoa(result))
	return nil
}

func (s Service) Reserve(clinicHistoryNumber, dniMedique int, date time.Time) error {
	begin, err := s.db.App().Begin()
	if err != nil {
		return err
	}

	// Dos transacciones no pueden reservar el mismo turno
	_, err = begin.Exec("set transaction isolation level serializable;")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	query := fmt.Sprintf("select reserve_appointment(%s, %s, '%s');", strconv.Itoa(clinicHistoryNumber), strconv.Itoa(dniMedique), kit.TimeToDateTimeString(date))

	var result bool
	err = begin.QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	begin.Commit()

	if result {
		fmt.Printf("Se ha reservado turno para la historia clinica %s y dniMedique %s a las %s \n", strconv.Itoa(clinicHistoryNumber), strconv.Itoa(dniMedique), date)
	} else {
		fmt.Printf("No se ha reservado turno para la historia clinica %s y dniMedique %s a las %s \n", strconv.Itoa(clinicHistoryNumber), strconv.Itoa(dniMedique), date)
	}

	return nil
}
