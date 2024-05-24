package main

import (
	"net"

	"github.com/charmbracelet/log"
	"github.com/d1agnoze/kitchen/services/auth/handler"
	"github.com/d1agnoze/kitchen/services/auth/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Info("Starting server")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	defer func() {
		log.Info("Stopping server")
		server.GracefulStop()
		lis.Close()
	}()

	pb.RegisterAuthServiceServer(server, &handler.AuthHandler{})
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
	log.Info("Server started on :50052")

}
