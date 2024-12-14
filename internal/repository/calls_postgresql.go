package repository

import (
	"call-center/internal/constants"
	"call-center/internal/domain"
	"errors"
	"gorm.io/gorm"
)

type CallPostgresqlRepo struct {
	db *gorm.DB
}

func NewCallPostgresqlRepo(db *gorm.DB) *CallPostgresqlRepo {
	return &CallPostgresqlRepo{
		db,
	}
}

// Создание заявки
func (r *CallPostgresqlRepo) Create(call *domain.Call) error {
	return r.db.Create(call).Error
}

// Получения всех заявок
func (r *CallPostgresqlRepo) GetAll() ([]*domain.Call, error) {
	var calls []*domain.Call
	err := r.db.Find(&calls).Error
	return calls, err
}

// Получения заявки по ID
func (r *CallPostgresqlRepo) GetByID(id uint64) (*domain.Call, error) {
	var call domain.Call
	result := r.db.First(&call, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, constants.ErrCallNotFound
		}
		return nil, result.Error
	}
	return &call, nil
}

// Обновления статуса заявки по ID
func (r *CallPostgresqlRepo) UpdateStatus(id uint64, status string) error {
	result := r.db.Model(&domain.Call{}).Where("id = ?", id).Update("status", status)
	if result.RowsAffected == 0 {
		return constants.ErrCallNotFound
	}
	return result.Error
}

// Удаления заявки по ID
func (r *CallPostgresqlRepo) DeleteById(id uint64) error {
	result := r.db.Delete(&domain.Call{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return constants.ErrCallNotFound
	}
	return nil
}
