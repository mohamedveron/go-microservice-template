package api

import (
	"context"
	pb "github.com/TheSocialLoanCompany/proto-buff-interface/loans_pb/api"
	"log"
)

func (s *Server) AddLoan(ctx context.Context, in *pb.NewLoan) (*pb.NewLoan, error) {
	log.Printf("Received: %v", in.GetName())

	loan := &pb.NewLoan{
		Id:       in.GetId(),
		Name:     in.GetName(),
		Amount:   in.GetAmount(),
		Duration: in.GetDuration(),
	}

	return loan, nil
}
