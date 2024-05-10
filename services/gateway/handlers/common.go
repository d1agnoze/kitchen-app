package handlers

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/d1agnoze/kitchen/services/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Connect() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to common service", err)
	}
	return conn
}

func Puck() (string, error) {
	conn := Connect()
	c := pb.NewCommonServiceClient(conn)
	defer conn.Close()

	resp, err := c.Puck(context.Background(), &pb.MessageDto{
		Name: `Ness`,
		Second: &pb.SecondDto{
			Count: 1,
		},
	})

	if err != nil {
		return "", err
	}
	return resp.Result, nil

}

func Ping() (string, error) {
	conn := Connect()
	c := pb.NewCommonServiceClient(conn)
	defer conn.Close()

	resp, err := c.Ping(context.Background(), &emptypb.Empty{})

	if err != nil {
		return "", err
	}
	return resp.Message, nil
}
