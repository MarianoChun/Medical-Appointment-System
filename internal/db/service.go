package db

import (
	_ "database/sql"
	"encoding/json"
	"github.com/boltdb/bolt"
	"gitlab.com/agustinesco/ruiz-escobar-mariano-tp/internal"
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
	tx, err := s.db.Bolt().Begin(true)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = s.syncPatients(tx)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	err = s.syncMedics(tx)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	err = s.syncConsultingRooms(tx)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	err = s.syncInsurances(tx)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	err = s.syncAppointments(tx)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return
	}
}

func (s Service) syncAppointments(tx *bolt.Tx) error {
	rows, err := s.db.App().Query("select nro_turno, fecha, nro_consultorio, dni_medique, nro_paciente, nro_obra_social_consulta, nro_afiliade_consulta, monto_paciente, monto_obra_social, f_reserva, estado from turno")
	if err != nil {
		log.Fatalln(err)
		return err
	}
	bucket, _ := tx.CreateBucketIfNotExists([]byte("appointments"))
	for rows.Next() {
		appointment := internal.Appointment{}
		err = rows.Scan(&appointment.Number, &appointment.Date, &appointment.ConsultingRoomNumber, &appointment.MedicDni, &appointment.PatientNumber, &appointment.PatientAmount, &appointment.InsuranceAmount, &appointment.ReserveDate, &appointment.Status)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		data, err := json.Marshal(appointment)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		err = bucket.Put([]byte(string(rune(appointment.Number))), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) syncConsultingRooms(tx *bolt.Tx) error {
	rows, err := s.db.App().Query("select nro_consultorio, nombre, domicilio, codigo_postal, telefono from consultorio")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	bucket, _ := tx.CreateBucketIfNotExists([]byte("consulting_rooms"))
	for rows.Next() {
		consultingRooms := internal.ConsultingRoom{}
		err = rows.Scan(&consultingRooms.Number, &consultingRooms.Name, &consultingRooms.Residence, &consultingRooms.PostalNumber, &consultingRooms.Phone)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		data, err := json.Marshal(consultingRooms)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		err = bucket.Put([]byte(string(rune(consultingRooms.Number))), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) syncMedics(tx *bolt.Tx) error {
	rows, err := s.db.App().Query("select dni_medique, nombre, apellido, especialidad, monto_consulta_privada, telefono from medique")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	bucket, _ := tx.CreateBucketIfNotExists([]byte("medics"))
	for rows.Next() {
		medic := internal.Medic{}
		err = rows.Scan(&medic.Dni, &medic.Name, &medic.Lastname, &medic.Specialty, &medic.AmountOfPrivateConsulting, &medic.Phone)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		data, err := json.Marshal(medic)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		err = bucket.Put([]byte(string(rune(medic.Dni))), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) syncPatients(tx *bolt.Tx) error {
	rows, err := s.db.App().Query("select nro_paciente, nombre, apellido, dni_paciente, f_nac, nro_obra_social, nro_afiliade, domicilio, telefono, email from paciente")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	bucket, _ := tx.CreateBucketIfNotExists([]byte("patients"))
	for rows.Next() {
		patient := internal.Patient{}
		err = rows.Scan(&patient.Number, &patient.Name, &patient.Lastname, &patient.Dni, &patient.Birthdate, &patient.InsuranceNumber, &patient.AffiliateNumber, &patient.Residence, &patient.Phone, &patient.Email)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		data, err := json.Marshal(patient)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		err = bucket.Put([]byte(string(rune(patient.Number))), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) syncInsurances(tx *bolt.Tx) error {
	rows, err := s.db.App().Query("select nro_obra_social, nombre, contacto_nombre, contacto_apellido, contacto_telefono, contacto_email from obra_social")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	bucket, _ := tx.CreateBucketIfNotExists([]byte("insurances"))
	for rows.Next() {
		insurance := internal.Insurance{}
		err = rows.Scan(&insurance.Number, &insurance.Name, &insurance.ContactName, &insurance.ContactLastname, &insurance.Phone, &insurance.Email)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		data, err := json.Marshal(insurance)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		err = bucket.Put([]byte(string(rune(insurance.Number))), data)
		if err != nil {
			return err
		}
	}
	return nil
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

func (s Service) ViewNoSQL() {
	buckets := []string{"appointments", "consulting_rooms", "medics", "patients", "insurances"}
	s.db.Bolt().View(func(tx *bolt.Tx) error {
		for i := 0; i < len(buckets); i++ {
			bucket := tx.Bucket([]byte(buckets[i]))

			bucket.ForEach(func(k, v []byte) error {
				log.Println(string(bucket.Get(k)))

				return nil
			})
		}
		return nil
	})
}
