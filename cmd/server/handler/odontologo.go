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

		o, err := oh.s.UpdateOdontologo(id, odontologo)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, o)
	}

}
func (oh *odontologoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Apellido  string `json:"apellido,omitempty"`
		Nombre    string `json:"nombre,omitempty"`
		Matricula string `json:"matricula" binding:"required"`
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
		_, err = oh.s.GetOdontologoById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("dentist not found"))
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		update := domain.Odontologo{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Matricula: r.Matricula,
		}
		//FIXME FALTA VALIDAR MATRICULA
		o, err := oh.s.UpdateOdontologo(id, update)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, o)
	}
}


func (oh *odontologoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		_, err = oh.s.GetOdontologoById(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("dentist not found"))
			return
		}
		err= oh.s.DeleteOdontologo(id)
		if err != nil {
			web.Failure(ctx,404,err)
			return
		}
		web.Success(ctx,204,"Message: Deleted")
	}
}