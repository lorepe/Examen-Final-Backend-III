package server

import (
	"Final/internal/domain"
	"Final/internal/paciente"
	"Final/pkg/web"
	"errors"
	"strconv"

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
	return func(ctx *gin.Context) {
		pacientes, err := ph.s.GetAll()
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, pacientes)
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
		p, err := ph.s.CreatePaciente(paciente)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, p)

	}

}

func (ph *pacienteHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		paciente, err := ph.s.GetPacienteById(id)
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, paciente)

	}
}
func (ph *pacienteHandler) Put() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		_, err = ph.s.GetPacienteById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("patient not found"))
			return
		}
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}

		var paciente domain.Paciente
		err = ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		p, err := ph.s.UpdatePaciente(id, paciente)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, p)
	}

}

func (ph *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Dni string `json:"dni" binding:"required"`
	}
	return func(ctx *gin.Context) {

		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}

		pacienteDb, err := ph.s.GetPacienteById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("patient not found"))
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		update := domain.Paciente{
			Nombre:    pacienteDb.Nombre,
			Apellido:  pacienteDb.Apellido,
			Domicilio: pacienteDb.Domicilio,
			Dni:       r.Dni,
			FechaAlta: pacienteDb.FechaAlta,
		}
		p, err := ph.s.UpdateDni(id, update)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, p)
	}
}

func (ph *pacienteHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		_, err = ph.s.GetPacienteById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("patient not found"))
			return
		}

		err = ph.s.DeletePaciente(id)
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, gin.H{"Success": "deleted"})
	}
}
