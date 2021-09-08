package grpcserver

import (
	"context"

	pb "github.com/n-arsenic/akunaki/server/generated/proto/echo"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedEchoServer

	helloMessage string
}

func (s *Server) Wellcome(ctx context.Context, req *pb.Request) (*pb.Responce, error) {
	return &pb.Responce{Message: s.helloMessage + ", " + req.GetUser()}, nil
}

func NewGRPCServer(msg string) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &Server{helloMessage: msg})

	return s
}
