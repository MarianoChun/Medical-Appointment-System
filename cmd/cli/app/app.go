package app

import (
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/fk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/instance"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/pk"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/sp/appointment"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/sp/insurance"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal/sync"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
)

type App struct {
	Database                     kit.Database
	Initializer                  instance.DatabaseInitializer
	PrimaryKeysCreator           pk.Creator
	PrimaryKeysDeleter           pk.Deleter
	ForeignKeysCreator           fk.Creator
	ForeignKeysDeleter           fk.Deleter
	DatabasesSynchronizer        sync.DatabasesSynchronizer
	AppointmentAttender          appointment.Attender
	AppointmentReserver          appointment.Reserver
	AppointmentCanceller         appointment.Canceller
	AppointmentGenerator         appointment.DateGenerator
	InsuranceSettlementGenerator insurance.SettlementGenerator
}

func NewApp() (App, error) {
	database, err := kit.NewDatabase()
	if err != nil {
		return App{}, err
	}

	return App{
		Database:                     database,
		Initializer:                  instance.NewDatabaseInitializer(database),
		PrimaryKeysCreator:           pk.NewPrimaryKeysCreator(database),
		PrimaryKeysDeleter:           pk.NewPrimaryKeysDeleter(database),
		ForeignKeysCreator:           fk.NewForeignKeysCreator(database),
		ForeignKeysDeleter:           fk.NewForeignKeysDeleter(database),
		DatabasesSynchronizer:        sync.NewDatabasesSynchronizer(database),
		AppointmentAttender:          appointment.NewAttender(database),
		AppointmentReserver:          appointment.NewReserver(database),
		AppointmentCanceller:         appointment.NewCanceller(database),
		AppointmentGenerator:         appointment.NewGenerator(database),
		InsuranceSettlementGenerator: insurance.NewSettlementGenerator(database),
	}, nil
}
