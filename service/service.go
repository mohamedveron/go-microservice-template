package service

import "github.com/TheSocialLoanCompany/go-microservice-template/persistence"

type Service struct {
	persistence *persistence.Persistence
}

func NewService(persistence *persistence.Persistence) *Service {

	return &Service{
		persistence: persistence,
	}
}
