package service

import (
	"context"
	"github.com/mohamedveron/go-microservice-template/domain"
	"github.com/mohamedveron/go-microservice-template/proxy/grpc/user_management"
	"github.com/mohamedveron/logger"
)

type UserServiceGrpcImplementation struct {
}


func (s *UserServiceGrpcImplementation) GetUserById(ctx context.Context, uuid string) (*domain.User, error) {
	userGrpcResponse, err := user_management.GetUserById(ctx, uuid)

	if err != nil {
		logger.Errorf("could not get user: %v", err)
		return nil, err
	}

	logger.Infof("user retrieved: %s", userGrpcResponse)

	return userGrpcResponse, nil
}