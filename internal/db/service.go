package db

import (
	_ "database/sql"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"log"
)

const (
	databaseScriptPath           = "sql/database.sql"
	schemaScriptPath             = "sql/schema.sql"
	pkScriptPath                 = "sql/pk/create.sql"
	fkScriptPath                 = "sql/fk/create.sql"
	dataFolderPath               = "sql/data"
	storedProceduresFolderPath   = "sql/sp"
	triggersProceduresFolderPath = "sql/triggers"

	initLogMessage       = "Instantiating Database"
	finishLogMessage     = "Database initialized!"
	errorOccurredMessage = "Error occurred"
)

type Service struct {
	db kit.Database
}

func NewService(db kit.Database) Service {
	return Service{
		db: db,
	}
}

func (s Service) SyncBetweenSQLAndNoSQL() {
}

func (s Service) Init() {
	log.Println(initLogMessage)

	err := kit.ExecuteScript(databaseScriptPath, s.db.Postgres())
	if err != nil {
		log.Fatalln(errorOccurredMessage, err)
		return
	}

	err = kit.ExecuteScript(schemaScriptPath, s.db.App())
	if err != nil {
		log.Fatalln(errorOccurredMessage, err)
		return
	}

	err = kit.ExecuteScript(pkScriptPath, s.db.App())
	if err != nil {
		log.Fatalln(errorOccurredMessage, err)
		return
	}

	err = kit.ExecuteScript(fkScriptPath, s.db.App())
	if err != nil {
		log.Fatalln(errorOccurredMessage, err)
		return
	}

	err = kit.ExecuteFunctionsCreation(storedProceduresFolderPath, s.db.App())
	if err != nil {
		log.Fatalln(errorOccurredMessage, err)
		return
	}

	err = kit.ExecuteFunctionsCreation(triggersProceduresFolderPath, s.db.App())
	if err != nil {
		log.Fatalln(errorOccurredMessage, err)
		return
	}

	err = kit.ExecuteScripts(dataFolderPath, s.db.App())
	if err != nil {
		log.Fatalln(errorOccurredMessage, err)
		return
	}

	log.Println(finishLogMessage)
}
