package kit

import (
	"database/sql"
	"fmt"
	"github.com/boltdb/bolt"
	_ "github.com/lib/pq"
	"log"
)

const (
	sqlDriverName     = "postgres"
	sqlUser           = "postgres"
	sqlPass           = "postgres"
	sqlHost           = "localhost"
	sqlPostgresDb     = "postgres"
	sqlApplicationDb  = "consultorio"
	sqlSslMode        = "disable"
	sqlDataSourceName = "user=%s password=%s host=%s dbname=%s sslmode=%s"

	boltPath = "consultorio.db"
	boltMode = 0600
)

type Database struct {
	postgresConnection *sql.DB
	appConnection      *sql.DB
	boltConnection     *bolt.DB
}

func NewDatabase() (Database, error) {
	postgresDataSource := fmt.Sprintf(sqlDataSourceName, sqlUser, sqlPass, sqlHost, sqlPostgresDb, sqlSslMode)
	postgresConnection, err := sql.Open(sqlDriverName, postgresDataSource)
	if err != nil {
		log.Fatal(fmt.Sprintf("can not connect with database %s", postgresDataSource), err)
		return Database{}, err
	}

	appDataSource := fmt.Sprintf(sqlDataSourceName, sqlUser, sqlPass, sqlHost, sqlApplicationDb, sqlSslMode)
	appConnection, err := sql.Open(sqlDriverName, appDataSource)
	if err != nil {
		log.Fatal(fmt.Sprintf("can not connect with database %s", appDataSource), err)
		return Database{}, err
	}

	boltConnection, err := bolt.Open(boltPath, boltMode, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("can not connect with database %s", boltConnection), err)
		return Database{}, err
	}

	return Database{
		postgresConnection: postgresConnection,
		appConnection:      appConnection,
		boltConnection:     boltConnection,
	}, nil
}

func (s Database) App() *sql.DB {
	return s.appConnection
}

func (s Database) Postgres() *sql.DB {
	return s.postgresConnection
}

func (s Database) Bolt() *bolt.DB {
	return s.boltConnection
}

func (s Database) Close() error {
	err := s.postgresConnection.Close()
	if err != nil {
		return err
	}

	err = s.appConnection.Close()
	if err != nil {
		return err
	}

	err = s.boltConnection.Close()
	if err != nil {
		return err
	}

	return nil
}
