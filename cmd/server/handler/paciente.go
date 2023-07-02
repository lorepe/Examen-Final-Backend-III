package server

import (
	_ "Final/docs"
	"Final/internal/domain"
	_ "Final/internal/domain"
	"Final/internal/paciente"
	"Final/pkg/web"
	_ "Final/pkg/web"
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

// ListPatients 	godoc
// @Summary 		List patients
// @Tags 			Paciente
// @Description 	Get patients list
// @Produce  		json
// @Success 		200 {object} []domain.Paciente
// @Router 			/pacientes [get]
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

// CreatePatient	godoc
// @Summary 		Create Patient
// @Tags 			Paciente
// @Description 	Post new patient
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Param 			odontologo body domain.Paciente{} true "Patient to store"
// @Success 		200 {object} domain.Paciente{}
// @Router 			/pacientes [post]
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


// FindById 			godoc
// @Summary				Get Single Patient by id.
// @Param				id path string true "get pacient by id"
// @Tags 				Paciente
// @Description			Return patient who matches idParam.
// @Produce				application/json
// @Success				200 {object} domain.Paciente{}
// @Router				/pacientes/{id} [get]
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

// UpdatePatient 	godoc
// @Summary 		Update Patient
// @Tags 			Paciente
// @Param 			paciente body domain.Paciente{} true "Patient to update"
// @Param			id path string true "id param"
// @Description 	Update patient
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Success 		200 {object} domain.Paciente{}
// @Router 			/pacientes/{id} [put]
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

// UpdateDNI 	godoc
// @Summary 		Update Patient-DNI
// @Tags 			Paciente
// @Param			id path string true "id param"
// @Param 			paciente body map[string]string true "DNI to patch"
// @Description 	Update DNI
// @Accept  		json
// @Produce  		json
// @Param 			token header string true "token"
// @Success 		200 {object} domain.Paciente{}
// @Router 			/pacientes/{id} [patch]
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

// DeletePatient		godoc
// @Summary			Delete Patient
// @Description		Remove Patient data by id.
// @Produce			application/json
// @Param			id path string true "id param"
// @Param 			token header string true "token"
// @Tags			Paciente
// @Success			200 {object} web.Response{}
// @Router			/pacientes/{id} [delete]
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
