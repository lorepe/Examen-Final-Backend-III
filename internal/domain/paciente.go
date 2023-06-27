package domain

type Paciente struct {
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Domicilio string `json:"domicilio" binding:"required"`
	Dni       string `json:"dni" binding:"required"`
	FechaAlta string `json:"fecha_alta" binding:"required"`
}
