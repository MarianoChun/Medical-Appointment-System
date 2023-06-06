package appointment

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
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

}

func (s Service) Cancel(dni int, dateFrom time.Time, dateTo time.Time) {

}

func (s Service) Generate(year int, month time.Month) {
}

func (s Service) Reserve(clinicHistoryNumber int, dni int, date time.Time) {
}
