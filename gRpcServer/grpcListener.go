package grpcserver

import (
	"fmt"
	"log"
	"net"
	"server-context/gRpcServer/protoGenerated"
	"time"

	"google.golang.org/grpc"
)

func RunServer() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		listener, err := net.Listen("tcp", ":9000")
		if err != nil {
			log.Fatalf("failed to listen on port 9000: %v\n", err)
		}

		grpcServer := grpc.NewServer()

		var fss protoGenerated.FileStorageServiceServer = &fileStorageServiceServer{fs: NewStorage()}
		protoGenerated.RegisterFileStorageServiceServer(grpcServer, fss)

		go func() {
			time.Sleep(5 * time.Second)
			grpcServer.GracefulStop()
		}()

		if err := grpcServer.Serve(listener); err != nil {
			fmt.Printf("failed to serve gRPC server on port 9000: %v\n", err)
		}

		done <- struct{}{}
	}()

	return done
}
