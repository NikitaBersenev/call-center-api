package v1

import (
	"call-center/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(router *gin.Engine) *gin.Engine {
	h.InitCallsRoutes(router)
	return router
}
