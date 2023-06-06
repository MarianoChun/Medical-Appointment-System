package appointment

import "gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"

type Attender struct {
	db kit.Database
}

func NewAttender(db kit.Database) Attender {
	return Attender{
		db: db,
	}
}

func (s Attender) Execute(appointmentNumber int) {

}
