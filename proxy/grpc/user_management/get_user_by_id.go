package user_management

import (
	"context"
	"github.com/mohamedveron/go-microservice-template/domain"
	"github.com/mohamedveron/logger"
	pb "github.com/mohamedveron/proto-buff-interface/user_pb/api"
)

func GetUserById(ctx context.Context, uuid string) (*domain.User, error) {

	// Contact the server and print out its response.

	grpc := NewUserManagementServiceClient("user-management-service-svc.onboarding.svc.cluster.local:50051")

	getUserRequest := &pb.GetUserRequest{Uuid: uuid}
	userGrpcResponse, err := grpc.Client.GetUser(ctx, getUserRequest)

	if err != nil {
		logger.Errorf("could not get user: %v", err)
		return nil, err
	}

	userData := domain.User{}
	/*err = userData.BuildUserFromGrpcResponse(userGrpcResponse)

	if err != nil {
		logger.Errorf("could not build user response: %v", err)
		return nil, err
	}*/

	logger.Infof("user retrieved: %s", userGrpcResponse.User)

	return &userData, nil
}
