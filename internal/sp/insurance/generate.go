package insurance

import "gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"

type SettlementGenerator struct {
	db kit.Database
}

func NewSettlementGenerator(db kit.Database) SettlementGenerator {
	return SettlementGenerator{
		db: db,
	}
}

func (s SettlementGenerator) Execute() {

}
