package handler

import (
	"context"

	"github.com/d1agnoze/kitchen/services/auth/pb"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
}

func (x *AuthHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{}, nil
}

func (x *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}

func (x *AuthHandler) Signup(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}
