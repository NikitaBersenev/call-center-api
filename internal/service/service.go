package service

import (
	"call-center/internal/domain"
	"github.com/go-playground/validator/v10"
)

type Services struct {
	Calls Calls
}

func NewServices(CallService *CallService) *Services {
	return &Services{
		Calls: CallService,
	}
}

type Calls interface {
	Create(call *domain.Call) error
	GetAll() ([]*domain.Call, error)
	GetByID(id uint64) (*domain.Call, error)
	UpdateStatus(id uint64, status string) error
	DeleteById(id uint64) error
	ValidatePhoneNumber(fl validator.FieldLevel) bool
}
