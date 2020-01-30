package main

import (
	"fmt"
	"log"
	"net"
	"os"

	grpcService "github.com/Penetration-Platform-Go/MongoDB-Service/grpc"
	mongodb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
)

func main() {
	var GRPCPORT = os.Getenv("GRPC_PORT")
	if len(GRPCPORT) == 0 {
		GRPCPORT = "8083"
	}
	var grpcPort = flag.StringP("grpc_port", "g", GRPCPORT, "Define the port where grpc service runs")
	flag.Parse()
	// start grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen grpc: %v", err)
	}
	log.Printf("Listening on: %s", *grpcPort)
	gs := grpc.NewServer()
	mongodb.RegisterMongoDBServer(gs, &grpcService.MongoDBService{})
	gs.Serve(lis)
}
