package repository

import (
	"call-center/internal/domain"
)

type Calls interface {
	Create(call *domain.Call) error
	GetAll() ([]*domain.Call, error)
	GetByID(id uint64) (*domain.Call, error)
	UpdateStatus(id uint64, status string) error
	DeleteById(id uint64) error
}
