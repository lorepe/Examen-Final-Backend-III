package server

import (
	"Final/internal/odontologo"

	"github.com/gin-gonic/gin"
)

type odontologoHandler struct {
	s odontologo.SeviceOdontologo
}

func NewOdontologoHandler(s odontologo.SeviceOdontologo) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

func (oh *odontologoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		//FIXME ERROR
		odontologos := oh.s.GetAll()
		c.JSON(200, odontologos)
		// web.Success(c, 200, odontologos)
	}
}
