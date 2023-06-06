package appointment

import (
	"fmt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"log"
	"strconv"
	"time"
)

type Service struct {
	db kit.Database
}

func NewService(db kit.Database) Service {
	return Service{
		db: db,
	}
}

func (s Service) Attend(appointmentNumber int) {
	query := fmt.Sprintf("select attend_appointment(%s);", strconv.Itoa(appointmentNumber))

	var result bool
	err := s.db.App().QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		return
	}

	if result {
		log.Printf("Turno %v atendido con exito. \n", appointmentNumber)
	} else {
		log.Printf("Ocurrió un error al atender el turno nro %v. \n", appointmentNumber)
	}
}

func (s Service) Cancel(dni int, dateFrom time.Time, dateTo time.Time) {
	query := fmt.Sprintf("select cancel_appointment(%s, '%s', '%s');", strconv.Itoa(dni), kit.TimeToDateString(dateFrom), kit.TimeToDateString(dateTo))

	var result int
	err := s.db.App().QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Los cantidad de turnos cancelados fueron %s \n", strconv.Itoa(result))
}

func (s Service) Generate(year int, month int) {
	query := fmt.Sprintf("select generate_appointments_in_month(%s, %s);", strconv.Itoa(year), strconv.Itoa(month))

	var result bool
	err := s.db.App().QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		return
	}

	if result {
		fmt.Printf("Generando turnos para el %v/%v. \n", month, year)
	} else {
		fmt.Printf("Los turnos diponibles para el %v/%v ya se encuentran generados. \n", year, month)
	}
}

func (s Service) Reserve(clinicHistoryNumber int, dni int, date time.Time) {
	query := fmt.Sprintf("select reserve_appointment(%s, %s, '%s');", strconv.Itoa(clinicHistoryNumber), strconv.Itoa(dni), kit.TimeToDateTimeString(date))

	var result bool
	err := s.db.App().QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		return
	}

	if result {
		fmt.Printf("Se ha reservado turno para la historia clinica %s y dni %s \n", strconv.Itoa(clinicHistoryNumber), strconv.Itoa(dni))
	} else {
		fmt.Printf("No se ha reservado turno para la historia clinica %s y dni %s \n", strconv.Itoa(clinicHistoryNumber), strconv.Itoa(dni))
	}
}
