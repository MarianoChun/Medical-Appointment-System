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
	begin, err := s.db.App().Begin()
	if err != nil {
		return err
	}

	// Con serializable nos aseguramos porque las transacciones se ejecutan secuencialmente y no enviaremos mail duplicados
	_, err = begin.Exec("set transaction isolation level serializable;")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	_, err = begin.Exec("select send_absence_emails();")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	begin.Commit()

	return nil
}

func (s Service) SendReminderEmails() error {
	begin, err := s.db.App().Begin()
	if err != nil {
		return err
	}

	// Con serializable nos aseguramos porque las transacciones se ejecutan secuencialmente y no enviaremos mail duplicados
	_, err = begin.Exec("set transaction isolation level serializable;")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	_, err = begin.Exec("select send_reminder_on_appointment_reserved();")
	if err != nil {
		log.Fatal(err)
		begin.Rollback()
		return err
	}

	begin.Commit()

	return nil
}
