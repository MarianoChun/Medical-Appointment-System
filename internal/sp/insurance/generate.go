package insurance

import "gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"

type Service struct {
	db kit.Database
}

func NewService(db kit.Database) Service {
	return Service{
		db: db,
	}
}

func (s Service) GenerateSettlements() {

}
