package main

import (
	"net"

	"github.com/charmbracelet/log"
	"github.com/d1agnoze/kitchen/services/common/handler"
	"github.com/d1agnoze/kitchen/services/common/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Info("Starting server")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	defer func() {
		log.Info("Stopping server")
		server.GracefulStop()
		lis.Close()
	}()

	pb.RegisterCommonServiceServer(server, &handler.CommonHandler{})
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
	log.Info("Server started on :50051")
}
