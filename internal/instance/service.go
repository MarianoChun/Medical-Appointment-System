package instance

import (
	_ "database/sql"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/kit"
	"log"
)

const (
	databaseScriptPath           = "sql/database.sql"
	schemaScriptPath             = "sql/schema.sql"
	dataFolderPath               = "sql/data"
	storedProceduresFolderPath   = "sql/sp"
	triggersProceduresFolderPath = "sql/triggers"

	initLogMessage       = "Instantiating Database"
	finishLogMessage     = "Database initialized!"
	errorOccurredMessage = "Error occurred"
)

type DatabaseInitializer struct {
	db kit.Database
}

func NewDatabaseInitializer(db kit.Database) DatabaseInitializer {
	return DatabaseInitializer{
		db: db,
	}
}

func (s DatabaseInitializer) Execute() {
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

	//TODO: Check SP initializing
	//err = kit.ExecuteScripts(storedProceduresFolderPath, s.db)
	//if err != nil {
	//	log.Fatalln(errorOccurredMessage, err)
	//	return
	//}

	err = kit.ExecuteScripts(triggersProceduresFolderPath, s.db.App())
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
