package main

import (
	"context"
	"fmt"
	hello "github.com/jsainthilaire/try-gripmock/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func run(log *log.Logger) error {
	log.Println("init GRPC server")
	defer log.Println("main: Completed")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	serverErrors := make(chan error, 1)

	s, err := NewServer(4000)
	if err != nil {
		return err
	}

	go func() {
		log.Printf("listening on port %d", 4000)
		serverErrors <- s.Serve()
	}()

	select {
	case err := <-serverErrors:
		return err
	case <-shutdown:
		log.Printf("stopping server")
		s.GracefulStop()
	}

	return nil
}

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

func (s *server) Hi(ctx context.Context, in *hello.HiRequest) (*hello.HiReply, error) {
	return &hello.HiReply{Reply: "hi " + in.GetName()}, nil
}

type Server interface {
	Serve() error
	GracefulStop()
	Hi(ctx context.Context, in *hello.HiRequest) (*hello.HiReply, error)
}

func NewServer(port int) (Server, error) {
	server := new(server)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return server, err
	}

	server.listener = listener

	server.grpcServer = grpc.NewServer()
	hello.RegisterGreeterServer(server.grpcServer, server)
	reflection.Register(server.grpcServer)
	return server, nil
}
