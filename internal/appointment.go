package internal

import "time"

type Appointment struct {
	Number               int       `json:"nro_turno"`
	Date                 time.Time `json:"fecha"`
	ConsultingRoomNumber int       `json:"nro_consultorio"`
	MedicDni             int       `json:"dni_medique"`
	PatientNumber        int       `json:"nro_paciente"`
	InsuranceNumber      int       `json:"nro_obra_social_consulta"`
	AffiliateNumber      int       `json:"nro_afiliade_consulta"`
	PatientAmount        float64   `json:"monto_paciente"`
	InsuranceAmount      float64   `json:"monto_obra_social"`
	ReserveDate          time.Time `json:"f_reserva"`
	Status               string    `json:"estado"`
}
