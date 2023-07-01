package server

import (
	"Final/internal/domain"
	"Final/internal/turno"
	"Final/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.ServiceTurno
}

func NewTurnoHandler(s turno.ServiceTurno) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}
func (th *turnoHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		turnos, err := th.s.GetAll()
		//FIXME pasar a reponse
		if err != nil {
			ctx.JSON(404, gin.H{"error": err})
			return
		}
		ctx.JSON(200, turnos)
	}
}

func (th *turnoHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turno domain.Turno
		err := ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		//TODO Validate empty
		t, err := th.s.CreateTurno(turno)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, t)

	}

}

func (th *turnoHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		turno, err := th.s.GetTurnoById(id)
		//FIXME ERROR ESTRUCTRA VACIA
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, turno)

	}
}

func (th *turnoHandler) Put() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		//FIXME Error id not exist
		_, err = th.s.GetTurnoById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("appointment not found"))
			return
		}
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}

		var turno domain.Turno
		err = ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		//TODO validate empty

		t, err := th.s.UpdateTurno(id, turno)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, t)
	}

}

func (th *turnoHandler) Patch() gin.HandlerFunc {
	//FIXME ALL
	type Request struct {
		Paciente    domain.Paciente   `json:"paciente,omitempty"`
		Odontologo  domain.Odontologo `json:"odontologo,omitempty"`
		FechaHora   string            `json:"fecha_hora" binding:"required"`
		Descripcion string            `json:"descripcion" binding:"required"`
	}
	return func(ctx *gin.Context) {

		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		//FIXME Error id not exist
		_, err = th.s.GetTurnoById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("appointment not found"))
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		//FIXME NO ADMITIR LA MATRICULA
		update := domain.Turno{
			Paciente:    r.Paciente,
			Odontologo:  r.Odontologo,
			FechaHora:   r.FechaHora,
			Descripcion: r.Descripcion,
		}
		//FIXME FALTA VALIDAR MATRICULA
		t, err := th.s.UpdateTurno(id, update)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, t)
	}
}

func (th *turnoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		_, err = th.s.GetTurnoById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("appointment not found"))
			return
		}
		err = th.s.DeleteTurno(id)
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 204, "Message: Deleted")
	}
}

func (th *turnoHandler) PostDNIMat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turno domain.TurnoAuxiliar
		err := ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		//TODO Validate empty
		err = th.s.CreateTurnoaAuxiliar(turno)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, turno)

	}

}
