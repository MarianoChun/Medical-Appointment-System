package insurance

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"log"
)

type Service struct {
	db kit.Database
}

func NewService(db kit.Database) Service {
	return Service{
		db: db,
	}
}

func (s Service) GenerateSettlements() {
	query := "select generate_insurance_settlements();"

	err := kit.ExecuteQuery(query, s.db.App())
	if err != nil {
		log.Fatal(err)
		return
	}
}
