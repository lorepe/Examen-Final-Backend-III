package server

import (
	"Final/internal/paciente"

	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.ServicePaciente
}

func NewPacienteHandler(s paciente.ServicePaciente) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}
func (ph *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		pacientes, err := ph.s.GetAll()
		if err != nil {
			c.JSON(404, gin.H{"error": "list not found"})
		}
		c.JSON(200, pacientes)
	}
}
