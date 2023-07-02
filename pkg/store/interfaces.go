package store

import "Final/internal/domain"

type StoreInterface interface {
	//MÉTODOS PARA OBTENER LISTADOS COMPLETOS
	GetAllPacientes() ([]domain.Paciente, error)
	GetAllTurnos() ([]domain.Turno, error)
	GetAllOdontologos() ([]domain.Odontologo, error)

	//MÉTODOS PARA CREAR ESTRUCTURAS
	PostOdontologo(domain.Odontologo) error
	PostPaciente(domain.Paciente) error
	PostTurno(domain.Turno) error

	//MÉTODOS PARA OBTENER ESTRUCTURAS POR ID
	GetOdontologoById(id int) (domain.Odontologo, error)
	GetPacienteById(id int) (domain.Paciente, error)
	GetTurnoById(id int) (domain.Turno, error)

	//MÉTODOS PARA ACTUALIZAR
	UpdateOdontologo(id int, o domain.Odontologo) error
	UpdatePaciente(id int, p domain.Paciente) error
	UpdateTurno(id int, t domain.Turno) error

	//MÉTODOS PARA ELIMINAR
	DeleteOdontologo(id int) error
	DeletePaciente(id int) error
	DeleteTurno(id int) error

	//MÉTODOS PARA TURNOS SOLICITADOS POR LÓGICA DE NEGOCIO
	PostTurnoDNIMat(ta domain.TurnoAuxiliar) error
	GetTurnosByDni(id int) ([]domain.Turno, error)

	//MÉTODOS PARA VERIFICAR DATOS

	//todo ver si sirve de algo
	VerificarMatricula(string) (bool, error)
}
