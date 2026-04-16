package main

import (
	"log"
	"micro-demo/user-service/internal/api"
	"micro-demo/user-service/internal/handler"
	"net/http"
)

func main() {
	h := handler.New()
	
	server, err := api.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server started on port :8080")
	http.ListenAndServe(":8080", server)
}
