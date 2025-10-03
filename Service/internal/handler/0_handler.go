package handler

import (
	"GreetService/internal/service"
)

type GreetHandler struct {
	GreetService *service.GreetService
}

func NewGreetHandler(service *service.GreetService) *GreetHandler {
	return &GreetHandler{GreetService: service}
}

