package domain

type Odontologo struct {
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}
