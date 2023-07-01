package server

import (
	"Final/internal/domain"
	"Final/internal/turno"
	"Final/pkg/web"
	"errors"

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