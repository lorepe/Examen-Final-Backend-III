package domain

type TurnoAuxiliar struct {
	PacienteDni   string    `json:"paciente_dni" binding:"required"`
	OdontologoMatricula string    `json:"odontologo_mat" binding:"required"`
	FechaHora    string `json:"fecha_hora" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`
}
