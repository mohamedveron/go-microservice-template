package api

import (
	"context"
	pb "github.com/TheSocialLoanCompany/proto-buff-interface/loans_pb/api"
)


func (s *Server) GetLoans(ctx context.Context, in *pb.LoanRequest) (*pb.LoanResponse, error) {

	var loans pb.LoanResponse

	return &loans, nil
}
