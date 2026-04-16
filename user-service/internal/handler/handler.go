package handler

import (
	"context"
	"micro-demo/user-service/internal/api"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	return &api.CreateUserResponse{
		Success: true,
	}, nil
}
