package app

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/db"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/fk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/pk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/sp/appointment"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/sp/insurance"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

type App struct {
	database           kit.Database
	DatabaseService    db.Service
	PrimaryKeysService pk.Service
	ForeignKeysService fk.Service
	Appointment        appointment.Service
	InsuranceService   insurance.Service
}

func NewApp() (App, error) {
	database, err := kit.NewDatabase()
	if err != nil {
		return App{}, err
	}

	return App{
		database:           database,
		DatabaseService:    db.NewService(database),
		PrimaryKeysService: pk.NewService(database),
		ForeignKeysService: fk.NewService(database),
		Appointment:        appointment.NewService(database),
		InsuranceService:   insurance.NewService(database),
	}, nil
}

func (s App) Close() {
	s.database.Close()
}
