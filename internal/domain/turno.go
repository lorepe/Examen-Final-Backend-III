package domain

type Turno struct {
	Id          int        `json:"id"`
	Paciente    Paciente   `json:"paciente" binding:"required"`
	Odontologo  Odontologo `json:"odontologo" binding:"required"`
	FechaHora   string     `json:"fecha_hora" binding:"required"`
	Descripcion string     `json:"descripcion" binding:"required"`
}
