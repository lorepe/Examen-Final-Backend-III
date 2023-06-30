package turno

import (
	"Final/internal/domain"
	"Final/pkg/store"
	"errors"
)

type RepositoryTurno interface {
	GetAll() ([]domain.Turno, error)
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) RepositoryTurno {
	return &repository{storage}
}

func (r *repository) GetAll() ([]domain.Turno, error) {
	turnos, err := r.storage.GetAllTurnos()
	if err != nil {
		return []domain.Turno{}, errors.New("list not found")
	}
	return turnos, nil
}
