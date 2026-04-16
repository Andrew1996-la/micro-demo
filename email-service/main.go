package main

import (
	"context"
	"log"
	emailpb "micro-demo/micro-demo/proto/email"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	emailpb.UnimplementedEmailServiceServer
}

func (s Server) SendEmail(ctx context.Context, req *emailpb.SendEmailRequest) (*emailpb.SendEmailResponse, error) {
	log.Println("отправляется сообщения для", req.To)
	log.Panicln("текст:", req.Message)

	return &emailpb.SendEmailResponse{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return
	}

	grpcServer := grpc.NewServer()

	emailpb.RegisterEmailServiceServer(grpcServer, &Server{})

	log.Println("Email service started on :50051")

	grpcServer.Serve(lis)
}
