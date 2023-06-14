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

func (s Service) GenerateSettlements() error {
	query := "select generate_insurance_settlements();"

	// Con serializable nos aseguramos que no se generen liquidaciones duplicadas ya que las transaccioens se ejecutan secuencialmente
	_, err := s.db.App().Exec("set transaction isolation level serializable;")
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = kit.ExecuteQuery(query, s.db.App())
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
