package service

import (
	"call-center/internal/constants"
	"call-center/internal/domain"
	"call-center/internal/repository"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"regexp"
)

type CallService struct {
	repo repository.Calls
}

func NewCallService(repo repository.Calls) *CallService {
	return &CallService{
		repo: repo,
	}
}

// Создание заявки
func (s *CallService) Create(call *domain.Call) error {
	if err := s.repo.Create(call); err != nil {
		return err
	}
	return nil
}

// Получения всех заявок
func (s *CallService) GetAll() ([]*domain.Call, error) {
	res, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Получения заявки по ID
func (s *CallService) GetByID(id uint64) (*domain.Call, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Обновления статуса заявки по ID
func (s *CallService) UpdateStatus(id uint64, status string) error {
	if err := s.repo.UpdateStatus(id, status); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constants.ErrCallNotFound
		}
		return err
	}
	return nil
}

// Удаления заявки по ID
func (s *CallService) DeleteById(id uint64) error {
	if err := s.repo.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// Валидация номера телефона
func (s *CallService) ValidatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber, ok := fl.Field().Interface().(string)
	if ok {
		regex := regexp.MustCompile(`^\+?[0-9-]*$`)
		return regex.MatchString(phoneNumber)
	}
	return true
}
