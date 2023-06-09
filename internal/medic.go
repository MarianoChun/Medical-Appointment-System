package internal

type Medic struct {
	Dni                       int     `json:"dni_medique"`
	Name                      string  `json:"nombre"`
	Lastname                  string  `json:"apellido"`
	Specialty                 string  `json:"especialidad"`
	AmountOfPrivateConsulting float64 `json:"monto_consulta_privada"`
	Phone                     string  `json:"telefono"`
}
