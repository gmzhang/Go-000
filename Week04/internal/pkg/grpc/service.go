package grpc

import (
	"google.golang.org/grpc"
	"context"
	"net"
	"log"
)

type Server struct {
	*grpc.Server

	address string
}

func NewServer(address string) *Server {
	svc := grpc.NewServer()
	return &Server{Server: svc, address: address}
}

func (s *Server) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.address)

	if err != nil {
		return err
	}

	log.Printf("grpc svc start: %s", s.address)

	go func() {
		<-ctx.Done()
		s.GracefulStop()
		log.Printf("grpc svc graceful stop")
	}()

	return s.Serve(l);
}
