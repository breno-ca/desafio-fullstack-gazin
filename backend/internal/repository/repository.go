package repository

import (
	"backend/internal/entity"
	"backend/internal/presenter"
	"backend/pkg/pagination"
)

type NivelRepository interface {
	Index(search, order, mode string, page pagination.Pagination) ([]presenter.Nivel, int, error)
	Create(nivel entity.Nivel) error
	Update(nivel entity.Nivel) error
	Delete(id string) error
	HasNivelAssociado(id string) (bool, error)
	CountDesenvolvedoresByNivel(id string) (int, error)
	IndexForSelectOptions() ([]presenter.NivelSelectOptions, int, error)
}

type DesenvolvedorRepository interface {
	Index(search, order, mode string, page pagination.Pagination) ([]presenter.Desenvolvedor, int, error)
	Create(desenvolvedor entity.Desenvolvedor) error
	Update(desenvolvedor entity.Desenvolvedor) error
	Delete(id string) error
	HasDesenvolvedorByID(id string) (bool, error)
}
