package server

import (
	"Final/internal/turno"

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
			ctx.JSON(404, gin.H{"error": "list not found"})
			return
		}
		ctx.JSON(200, turnos)
	}
}
