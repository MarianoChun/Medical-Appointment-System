package kit

import (
	"database/sql"
	"log"
	strings "strings"
)

const (
	executingScriptsMessage = "Executing scripts of folder %s\n"
	executingScriptMessage  = "Executing %s script\n"
)

func ExecuteFunctionsCreation(folderPath string, db *sql.DB) error {
	log.Printf(executingScriptsMessage, folderPath)

	files, err := GetFilesPathByFolder(folderPath)
	if err != nil {
		return err
	}

	for i := 0; i < len(files); i++ {
		file := files[i]

		query, err := ReadFile(file)
		if err != nil {
			return err
		}

		err = ExecuteQuery(query, db)
		if err != nil {
			return err
		}
	}

	return nil
}

func ExecuteScripts(folderPath string, db *sql.DB) error {
	log.Printf(executingScriptsMessage, folderPath)

	files, err := GetFilesPathByFolder(folderPath)
	if err != nil {
		return err
	}

	for i := 0; i < len(files); i++ {
		err = ExecuteScript(files[i], db)
		if err != nil {
			return err
		}
	}
	return nil
}

func ExecuteScript(path string, db *sql.DB) error {
	log.Printf(executingScriptMessage, path)

	script, err := ReadFile(path)
	if err != nil {
		return err
	}

	queries := strings.Split(script, ";")
	for i := 0; i < len(queries); i++ {
		err = ExecuteQuery(queries[i], db)
		if err != nil {
			return err
		}
	}

	return nil
}

func ExecuteQuery(query string, db *sql.DB) error {
	if strings.EqualFold(query, "") {
		return nil
	}

	query = strings.Replace(query, "\n", "", -1)

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
