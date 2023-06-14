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
	begin, err := s.db.App().Begin()
	if err != nil {
		return err
	}

	// Con serializable nos aseguramos que no se generen liquidaciones duplicadas ya que las transaccioens se ejecutan secuencialmente
	_, err = begin.Exec("set transaction isolation level serializable;")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	_, err = begin.Exec("select generate_insurance_settlements();")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	begin.Commit()

	return nil
}
