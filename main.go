package main

import (
	"github.com/mohamedveron/go-microservice-template/api"
	"github.com/mohamedveron/go-microservice-template/persistence"
	"github.com/mohamedveron/go-microservice-template/service"
	"github.com/mohamedveron/logger"
	pb "github.com/mohamedveron/proto-buff-interface/loans_pb/api"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":50051"
)

func main() {

	// db configuration
	dbSettings := persistence.DBSettings{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "123456",
		DbName:   "tslc",
	}

	//initialize the service layers
	persistenceLayer := persistence.NewPersistence(&dbSettings)
	serviceLayer := service.NewService(persistenceLayer)

	server := api.NewServer(serviceLayer)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoanServer(s, server)
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

}