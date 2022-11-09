package user_management

import (
	"github.com/TheSocialLoanCompany/logger"
	pb "github.com/TheSocialLoanCompany/proto-buff-interface/user_pb/api"
	"google.golang.org/grpc"
)

type UserManagementServiceClient struct {
	Client pb.UserServiceClient
}

func NewUserManagementServiceClient(address string)  UserManagementServiceClient{

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}

	//defer conn.Close()

	return UserManagementServiceClient{
		Client: pb.NewUserServiceClient(conn),
	}

}