package internal

import "time"

type Patient struct {
	Number          int       `json:"nro_paciente"`
	Name            string    `json:"nombre"`
	Lastname        string    `json:"apellido"`
	Dni             int       `json:"dni_paciente"`
	Birthdate       time.Time `json:"f_nac"`
	InsuranceNumber int       `json:"nro_obra_social"`
	AffiliateNumber int       `json:"nro_afiliade"`
	Residence       string    `json:"domicilio"`
	Phone           string    `json:"telefono"`
	Email           string    `json:"email"`
}
