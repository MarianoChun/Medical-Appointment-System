package email

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

func (s Service) SendAbsenseEmails() error {
	query := "select send_absence_emails();"

	// Con serializable nos aseguramos porque las transacciones se ejecutan secuencialmente y no enviaremos mail duplicados
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

func (s Service) SendReminderEmails() error {
	query := "select send_reminder_on_appointment_reserved();"

	// Con serializable nos aseguramos porque las transacciones se ejecutan secuencialmente y no enviaremos mail duplicados
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
