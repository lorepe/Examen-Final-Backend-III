package server

import (
	"Final/internal/domain"
	"Final/internal/odontologo"
	"Final/pkg/web"
	"errors"
	"strconv"

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

func (oh *odontologoHandler) Post() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var odontologo domain.Odontologo

		err := ctx.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		//TODO Validate empty
		o, err := oh.s.CreateOdontologo(odontologo)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, o)

	}

}

func (oh *odontologoHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		odontologo, err := oh.s.GetOdontologoById(id)
		//FIXME ERROR ESTRUCTRA VACIA
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, odontologo)

	}
}

func (oh *odontologoHandler) Put() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		//FIXME Error id not exist
		_, err = oh.s.GetOdontologoById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("dentist not found"))
			return
		}
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}

		var odontologo domain.Odontologo
		err = ctx.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		//TODO validate empty

		o, err:= oh.s.UpdateOdontologo(id,odontologo)
		if err != nil {
			web.Failure(ctx,409,err)
			return 
		}
		web.Success(ctx,200,o)
	}

}
