package server

import (
	"Final/internal/domain"
	"Final/internal/odontologo"
	"Final/pkg/web"
	"errors"

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
	return func(ctx *gin.Context) {
		odontologos, err := oh.s.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{"error": "list not found"})
		}
		ctx.JSON(200, odontologos)
		// web.Success(c, 200, odontologos)
	}
}

func (oh *odontologoHandler)Post() gin.HandlerFunc  {

	return func(ctx *gin.Context) {
		var odontologo domain.Odontologo
		
		err := ctx.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		//TODO Validate empty
		o, err:= oh.s.CreateOdontologo(odontologo)
		if err != nil {
			web.Failure(ctx , 400, err)
			return 
		}
		web.Success(ctx,201,o)

	}

}