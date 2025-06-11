package main

import (
	"log"

	"github.com/rivaldoyoseps/schedule-service/api/proto_gen/schedulepb"
	"github.com/rivaldoyoseps/schedule-service/internal/config"
	grpcHandler "github.com/rivaldoyoseps/schedule-service/internal/delivery/grpc"
	"github.com/rivaldoyoseps/schedule-service/internal/repository"
	"github.com/rivaldoyoseps/schedule-service/internal/usecase"
	"google.golang.org/grpc"
)

func main() {
	db:= config.NewGormConfig()
	repo := repository.NewScheduleRepository(db)
	scheduleUC := usecase.NewScheduleUseCase(repo)

	// Jalankan gRPC server
	go func() {
		list := config.NewGrpcListener()
		grpcSrvHandler := grpcHandler.NewScheduleGRPCHandler(scheduleUC)
		grpcServer := grpc.NewServer()
		schedulepb.RegisterScheduleServiceServer(grpcServer, grpcSrvHandler)
		log.Println("gRPC server running on :50051")
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("Gagal menjalankan gRPC server: %v", err)
		}
	}()

	// Jalankan REST API Fiber
	httpServer := config.NewFiberConfig()
	config.StartHttpServer(&config.HttpServerConfig{
		DB: db,
		APP: httpServer,
	})
}
