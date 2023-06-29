package server

import (
	"Final/internal/domain"
	"Final/internal/paciente"
	"Final/pkg/web"
	"errors"

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
func (ph *pacienteHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paciente domain.Paciente
		err := ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		//TODO Validate empty
		p, err := ph.s.CreatePaciente(paciente)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, p)

	}

}
