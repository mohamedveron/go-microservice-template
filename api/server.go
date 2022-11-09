package api

import (
	"github.com/mohamedveron/go-microservice-template/service"
	pb "github.com/mohamedveron/proto-buff-interface/loans_pb/api"
)

type Server struct {
	pb.UnimplementedLoanServer
	service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{
		UnimplementedLoanServer: pb.UnimplementedLoanServer{},
		service: service,
	}
}
