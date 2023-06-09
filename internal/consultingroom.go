package internal

type ConsultingRoom struct {
	Number       int    `json:"nro_consultorio"`
	Name         string `json:"nombre"`
	Residence    string `json:"domicilio"`
	PostalNumber string `json:"codigo_postal"`
	Phone        string `json:"phone"`
}
