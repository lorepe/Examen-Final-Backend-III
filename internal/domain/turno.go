package domain

type Turno struct {
	Id int `json:"id"`
	// Paciente    Paciente   `json:"paciente" binding:"required"`
	PacienteID int `json:"paciente_id" binding:"required"`
	// Odontologo  Odontologo `json:"odontologo" binding:"required"`
	OdontologoID int    `json:"odontologo_id" binding:"required"`
	FechaHora    string `json:"fecha_hora" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`
}
