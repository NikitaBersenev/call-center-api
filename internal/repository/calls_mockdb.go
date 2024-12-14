package repository

import (
	"call-center/internal/domain"
	"gorm.io/gorm"
)

type CallMockRepo struct {
	calls  map[uint64]*domain.Call
	nextID uint64
}

func NewCallMockRepo() *CallMockRepo {
	return &CallMockRepo{
		calls:  make(map[uint64]*domain.Call),
		nextID: 1,
	}
}

// Создание заявки
func (r *CallMockRepo) Create(call *domain.Call) error {
	call.ID = r.nextID
	r.calls[r.nextID] = call
	r.nextID++
	return nil
}

// Получения всех заявок
func (r *CallMockRepo) GetAll() ([]*domain.Call, error) {
	var result []*domain.Call
	for _, call := range r.calls {
		result = append(result, call)
	}
	return result, nil
}

// Получения заявки по ID
func (r *CallMockRepo) GetByID(id uint64) (*domain.Call, error) {
	call, exists := r.calls[id]
	if !exists {
		return nil, gorm.ErrRecordNotFound
	}
	return call, nil
}

// Обновления статуса заявки по ID
func (r *CallMockRepo) UpdateStatus(id uint64, status string) error {
	call, exists := r.calls[id]
	if !exists {
		return gorm.ErrRecordNotFound
	}
	call.Status = status
	return nil
}

// Удаления заявки по ID
func (r *CallMockRepo) DeleteById(id uint64) error {
	_, exists := r.calls[id]
	if !exists {
		return gorm.ErrRecordNotFound
	}
	delete(r.calls, id)
	return nil
}
