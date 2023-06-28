package store

import "Final/internal/domain"

type StoreInterface interface {
	GetAllOdontologos() ([]domain.Odontologo, error)
	GetAllPacientes() ([]domain.Paciente, error)
	GetAllTurnos() ([]domain.Turno, error)

	PostOdontologo(domain.Odontologo) error
}
