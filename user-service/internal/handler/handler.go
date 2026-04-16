package handler

import (
	"context"
	"fmt"
	emailpb "micro-demo/micro-demo/proto/email"
	"micro-demo/user-service/internal/api"
)

type Handler struct {
	emailClient emailpb.EmailServiceClient
}

func New(client emailpb.EmailServiceClient) *Handler {
	return &Handler{
		emailClient: client,
	}
}

func (h *Handler) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	checkRes, err := h.emailClient.CheckEmail(ctx, &emailpb.CheckEmailRequest{
		Email: req.Email,
	})

	if err != nil {
		return nil, err
	}

	if checkRes.Exist {
		return nil, fmt.Errorf("email already exist")
	}

	_, err = h.emailClient.SendEmail(ctx, &emailpb.SendEmailRequest{
		To:      req.Email,
		Message: "hello! how are you?",
	})

	if err != nil {
		return nil, err
	}

	return &api.CreateUserResponse{
		Success: true,
	}, nil
}

func (h *Handler) Ping(ctx context.Context) (*api.PingResponse, error) {
	return &api.PingResponse{
		Messages: "pong",
	}, nil
}
