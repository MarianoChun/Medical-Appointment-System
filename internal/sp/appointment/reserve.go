package appointment

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"time"
)

type Reserver struct {
	db kit.Database
}

func NewReserver(db kit.Database) Reserver {
	return Reserver{
		db: db,
	}
}

func (s Reserver) Execute(clinicHistoryNumber int, dni int, date time.Time) {
}
