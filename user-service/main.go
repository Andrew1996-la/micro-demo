package main

import (
	"log"
	emailpb "micro-demo/micro-demo/proto/email"
	"micro-demo/user-service/internal/api"
	"micro-demo/user-service/internal/handler"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatal(err)
	}

	client := emailpb.NewEmailServiceClient(conn)

	h := handler.New(client)

	server, err := api.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server started on port :8080")
	http.ListenAndServe(":8080", server)
}
