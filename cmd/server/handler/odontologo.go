package server

import (
	"Final/internal/domain"
	"fmt"

	"Final/internal/odontologo"
	"Final/pkg/web"

	"errors"
	"net/http"
	_ "net/http"
	"strconv"

	_ "Final/docs"

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

// ListDentist godoc
// @Summary 		List dentists
// @Tags 			Odontologo
// @Description 	Get dentist list
// @Produce  		json
// @Success 		200 {object} []domain.Odontologo
// @Router 			/odontologos [get]
func (oh *odontologoHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		odontologos, err := oh.s.GetAll()
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, odontologos)
	}
}

// CreateDentist 	godoc
// @Summary 		Create Dentist
// @Tags 			Odontologo
// @Description 	Post new dentist
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Param 			odontologo body domain.Odontologo{} true "Dentist to store"
// @Success 		200 {object} domain.Odontologo{}
// @Router 			/odontologos [post]
func (oh *odontologoHandler) Post() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var odontologo domain.Odontologo

		err := ctx.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid json"))
			return
		}
		o, err := oh.s.CreateOdontologo(odontologo)
		if err != nil {
			web.Failure(ctx,http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, o)

	}

}

// FindById 			godoc
// @Summary				Get Single Dentist by id.
// @Param				id path string true "get dentist by id"
// @Tags 				Odontologo
// @Description			Return dentist who matches idParam.
// @Produce				application/json
// @Success				200 {object} domain.Odontologo{}
// @Router				/odontologos/{id} [get]
func (oh *odontologoHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx,http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		odontologo, err := oh.s.GetOdontologoById(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx,http.StatusOK, odontologo)

	}
}

// UpdateDentist 	godoc
// @Summary 		Update Dentist
// @Tags 			Odontologo
// @Param 			odontologo body domain.Odontologo{} true "Dentist to update"
// @Param			id path string true "id param"
// @Description 	Update dentist
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Success 		200 {object} domain.Odontologo{}
// @Router 			/odontologos/{id} [put]
func (oh *odontologoHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		var odontologo domain.Odontologo
		err = ctx.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid json"))
			return
		}
		o, err := oh.s.UpdateOdontologo(id, odontologo)

		if fmt.Sprint(err) == "dentist not found" {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		} else if err != nil {
			web.Failure(ctx,http.StatusConflict, err)
			return
		}
		web.Success(ctx,http.StatusOK, o)
	}

}

// METODO DE NEGOCIO PARA ACTUALIZAR MATRICULA

// UpdateDentist 	godoc
// @Summary 		Update Dentist-Registration
// @Tags 			Odontologo
// @Param			id path string true "id param"
// @Param 			odontologo body map[string]string true "Registration to patch"
// @Description 	Update registration
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Success 		200 {object} domain.Odontologo{}
// @Router 			/odontologos/{id} [patch]
func (oh *odontologoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Matricula string `json:"matricula" binding:"required"`
	}
	return func(ctx *gin.Context) {

		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx,http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid json"))
			return
		}
		o, err := oh.s.UpdateMatricula(id, r.Matricula)
		if fmt.Sprint(err) == "dentist not found" {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		} else if err != nil {
			web.Failure(ctx, http.StatusConflict, err)
			return
		}
		web.Success(ctx, http.StatusOK, o)
	}
}

// DeleteDentist		godoc
// @Summary			Delete Dentist
// @Description		Remove Dentist data by id.
// @Produce			application/json
// @Param			id path string true "id param"
// @Param 			token header string true "token"
// @Tags			Odontologo
// @Success			200 {object} web.Response{}
// @Router			/odontologos/{id} [delete]
func (oh *odontologoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx,http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = oh.s.DeleteOdontologo(id)
		if err != nil {
			web.Failure(ctx,http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, gin.H{"Success": "deleted"})
	}
}
