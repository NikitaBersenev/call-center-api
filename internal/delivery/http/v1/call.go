package v1

import (
	"call-center/internal/constants"
	"call-center/internal/domain"
	"call-center/pkg/myerr"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) InitCallsRoutes(router *gin.Engine) *gin.Engine {
	calls := router.Group("/calls")
	calls.POST("", h.CreateCall)
	calls.GET("", h.GetCalls)
	calls.GET("/:id", h.GetCallByID)
	calls.PATCH("/:id/status", h.UpdateCallStatus)
	calls.DELETE("/:id", h.DeleteCall)
	return router
}

// Обработка создания новой заявки
func (h *Handler) CreateCall(c *gin.Context) {
	var request domain.CallCreateRequest
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.ValidatorErrorStr(err)})
		return
	}

	call := &domain.Call{
		ClientName:  request.ClientName,
		PhoneNumber: request.PhoneNumber,
		Description: request.Description,
		Status:      "открыта",
		CreatedAt:   time.Now(),
	}

	if err := h.services.Calls.Create(call); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, call)
}

// Обработка получения всех заявок
func (h *Handler) GetCalls(c *gin.Context) {
	calls, err := h.services.Calls.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, calls)
}

// Обработка получения заявки по ID
func (h *Handler) GetCallByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrCallInvalidId})
		return
	}

	call, err := h.services.Calls.GetByID(id)
	if err != nil {
		var statusResponse = http.StatusInternalServerError
		if errors.Is(err, constants.ErrCallNotFound) {
			statusResponse = http.StatusNotFound
		}
		c.JSON(statusResponse, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, call)
}

// Обработка обновления статуса заявки
func (h *Handler) UpdateCallStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrCallInvalidId})
		return
	}

	var request struct {
		Status string `json:"status" binding:"required,oneof=открыта закрыта"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.services.Calls.UpdateStatus(id, request.Status); err != nil {
		var statusResponse = http.StatusInternalServerError
		if errors.Is(err, constants.ErrCallNotFound) {
			statusResponse = http.StatusNotFound
		}
		c.JSON(statusResponse, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// Обработка удаления заявки
func (h *Handler) DeleteCall(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrCallInvalidId})
		return
	}

	if err := h.services.Calls.DeleteById(id); err != nil {
		var statusResponse = http.StatusInternalServerError
		if errors.Is(err, constants.ErrCallNotFound) {
			statusResponse = http.StatusNotFound
		}
		c.JSON(statusResponse, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
