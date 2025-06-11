package config

import (
	"log"
	"net"
)

func NewGrpcListener() net.Listener {
	grpcPort := ":50053"

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Gagal listen grpcPort %s: %v", grpcPort, err)
	}
	log.Printf("gRPC server listening on %s", grpcPort)
	return lis
}
