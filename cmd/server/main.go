package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(fmt.Errorf("error running grpc server: %w", err))
	}
}

func run() error {
	listener, err := net.Listen("tcp", "8000")
	if err != nil {
		fmt.Println(fmt.Errorf("listener error: %w", err))
	}

	server := grpc.NewServer()
	defer server.GracefulStop()

	server.Serve(listener)
	return nil
}
