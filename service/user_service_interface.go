package service

import (
	"context"
	"github.com/TheSocialLoanCompany/go-microservice-template/domain"
)

type UserServiceInterface interface {
	GetUserById(ctx context.Context, uuid string) (*domain.User, error)
}
