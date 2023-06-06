package appointment

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"time"
)

type Canceller struct {
	db kit.Database
}

func NewCanceller(db kit.Database) Canceller {
	return Canceller{
		db: db,
	}
}

func (s Canceller) Execute(dni int, dateFrom time.Time, dateTo time.Time) {

}
