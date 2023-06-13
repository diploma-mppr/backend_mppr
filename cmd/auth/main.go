package main

import (
	"fmt"
	GRPCAuthHandler "gitgub.com/diploma-mppr/backend_mppr/internal/microservices/auth/AuthHandler"
	"gitgub.com/diploma-mppr/backend_mppr/internal/microservices/auth/AuthPB"
	"gitgub.com/diploma-mppr/backend_mppr/internal/microservices/auth/AuthRepository"
	"gitgub.com/diploma-mppr/backend_mppr/internal/microservices/auth/AuthUseCase"
	"gitgub.com/diploma-mppr/backend_mppr/tools"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	pgxManager, err := tools.NewPostgresqlX()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	authRepo := AuthRepository.NewRepositoryAuth(pgxManager)
	authUcase := AuthUseCase.NewUseCaseAuth(authRepo)
	authHandler := GRPCAuthHandler.NewHandlerAuth(authUcase)

	authpb.RegisterAuthMicroserviceServer(server, authHandler)

	fmt.Println("starting server at :8001")
	server.Serve(lis)
}
