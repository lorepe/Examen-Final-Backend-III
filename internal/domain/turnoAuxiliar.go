package domain

type TurnoAuxiliar struct {
	//FIXME CAMBIAR A DNI y MAT

	PacienteId   int    `json:"paciente_id" binding:"required"`
	OdontologoId int    `json:"odontologo_id" binding:"required"`
	FechaHora    string `json:"fecha_hora" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`
}
