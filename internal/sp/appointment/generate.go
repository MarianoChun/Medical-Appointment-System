package appointment

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"time"
)

type DateGenerator struct {
	db kit.Database
}

func NewGenerator(db kit.Database) DateGenerator {
	return DateGenerator{
		db: db,
	}
}

func (s DateGenerator) Execute(year int, month time.Month) {
}
