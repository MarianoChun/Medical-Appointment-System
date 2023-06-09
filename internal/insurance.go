package internal

type Insurance struct {
	Number          int    `json:"nro_obra_social"`
	Name            string `json:"nombre"`
	ContactName     string `json:"contacto_nombre"`
	ContactLastname string `json:"contacto_apellido"`
	Phone           string `json:"contacto_telefono"`
	Email           string `json:"contacto_email"`
}
