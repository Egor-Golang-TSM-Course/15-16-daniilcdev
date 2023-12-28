package grpcserver

import "fmt"

func StartServer() {
	fmt.Println("starting gRPC server on port 9000...")

	done := RunServer()

	<-done

	fmt.Println("gRPC server on port 9000 exit")
}
