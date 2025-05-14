package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LavaJover/shvark-user-service/internal/config"
	"github.com/LavaJover/shvark-user-service/internal/delivery/grpcapi"
	"github.com/LavaJover/shvark-user-service/internal/infrastructure/postgres"
	"github.com/LavaJover/shvark-user-service/internal/usecase"
	userpb "github.com/LavaJover/shvark-user-service/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	// Reading config
	cfg := config.MustLoad()

	// Init database
	db := postgres.MustInitDB(cfg.Dsn)

	// Init user repo
	userRepo, err := postgres.NewUserRepository(db)
	if err != nil{
		log.Fatalf("failed to init user repository: %v\n", err.Error())
	}

	// Init user usecase
	uc := usecase.NewUserUsecase(userRepo)

	// Creating gRPC server
	grpcServer := grpc.NewServer()
	userHandler := grpcapi.UserHandler{UserUsecase: uc}

	userpb.RegisterUserServiceServer(grpcServer, &userHandler)

	// Start
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC server started on %s:%s\n", cfg.Host, cfg.Port)
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v\n", err)
	}
}