package http

import (
	v1 "call-center/internal/delivery/http/v1"
	"call-center/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phonenumber", h.services.Calls.ValidatePhoneNumber)
	}

	handlerV1 := v1.NewHandler(h.services)
	handlerV1.Init(router)
	return router
}
