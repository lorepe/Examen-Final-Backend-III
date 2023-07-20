package server

import (
	"Final/internal/domain"
	"Final/internal/turno"
	"Final/pkg/web"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "Final/docs"
	_ "net/http"
)

type turnoHandler struct {
	s turno.ServiceTurno
}

func NewTurnoHandler(s turno.ServiceTurno) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// ListAppointment godoc
// @Summary 		List appointments
// @Tags 			Turno
// @Description 	Get appointment list
// @Produce  		json
// @Success 		200 {object} []domain.Turno
// @Router 			/turnos [get]
func (th *turnoHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		turnos, err := th.s.GetAll()
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, turnos)
	}
}

// CreateAppointment 	godoc
// @Summary 		Create Appointment
// @Tags 			Turno
// @Description 	Post new appointment
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Param 			turno body domain.Turno{} true "Appointment to store"
// @Success 		200 {object} domain.Turno{}
// @Router 			/turnos [post]
func (th *turnoHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turno domain.Turno
		err := ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		t, err := th.s.CreateTurno(turno)
		if err != nil {
			web.Failure(ctx, 500, err)
			return
		}
		web.Success(ctx, 201, t)

	}

}

// FindById 			godoc
// @Summary				Get Single Appointment by id.
// @Param				id path string true "get appointment by id"
// @Tags 				Turno
// @Description			Return appointment who matches idParam.
// @Produce				application/json
// @Success				200 {object} domain.Turno{}
// @Router				/turnos/{id} [get]
func (th *turnoHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		turno, err := th.s.GetTurnoById(id)
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, turno)

	}
}

// UpdateAppointment 	godoc
// @Summary 		Update Appointment
// @Tags 			Turno
// @Param 			turno body domain.Turno{} true "Apponitment to update"
// @Param			id path string true "id param"
// @Description 	Update appointment
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Success 		200 {object} domain.Turno{}
// @Router 			/turnos/{id} [put]
func (th *turnoHandler) Put() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		var turno domain.Turno
		err = ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		t, err := th.s.UpdateTurno(id, turno)
		if fmt.Sprint(err) == "appointment not found" {
			web.Failure(ctx, 404, err)
			return
		} else if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, t)
	}

}

// UpdateAppointment 	godoc
// @Summary 		Update Appointment-Date
// @Tags 			Turno
// @Param			id path string true "id param"
// @Param 			turno body map[string]string true "Date to patch"
// @Description 	Update date
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Success 		200 {object} domain.Turno{}
// @Router 			/turnos/{id} [patch]
func (th *turnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		FechaHora string `json:"fecha_hora" binding:"required"`
	}
	return func(ctx *gin.Context) {

		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		t, err := th.s.UpdateCitaTurno(id, r.FechaHora)
		if fmt.Sprint(err) == "appointment not found" {
			web.Failure(ctx, 404, err)
			return
		} else if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, t)
	}
}

// DeleteAppointment		godoc
// @Summary			Delete Appointment
// @Description		Remove Appointment data by id.
// @Produce			application/json
// @Param			id path string true "id param"
// @Param 			token header string true "token"
// @Tags			Turno
// @Success			200 {object} web.Response{}
// @Router			/turnos/{id} [delete]
func (th *turnoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		err = th.s.DeleteTurno(id)
		if fmt.Sprint(err) == "appointment not found" {
			web.Failure(ctx, 404, err)
			return
		} else if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		web.Success(ctx, 200, gin.H{"Success": "deleted"})
	}
}

// CreateAppointment 	godoc
// @Summary 		Create Appointment with dni and registration
// @Tags 			Turno
// @Description 	Post new appointment with dni and registration
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Param 			turno body domain.TurnoAuxiliar{} true "Appointment to store"
// @Success 		200 {object} domain.Turno{}
// @Router 			/turnos/turnoauxiliar [post]
func (th *turnoHandler) PostDNIMat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turno domain.TurnoAuxiliar
		err := ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		err = th.s.CreateTurnoaAuxiliar(turno)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, turno)

	}

}

// ListAppointmentbyDni godoc
// @Summary 		List appointments by patient DNI
// @Tags 			Turno
// @Description 	Get appointment list by patient DNI
// @Param			dni query string true "dni param"
// @Produce  		json
// @Success 		200 {object} []domain.Turno
// @Router 			/turnos/paciente [get]
func (th *turnoHandler) GetAllByDni() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dniParam := ctx.Query("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid dni"))
			return
		}
		turnos, err := th.s.GetAllByDni(dni)
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}
		web.Success(ctx, 200, turnos)
	}
}
