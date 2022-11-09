package service

import "github.com/mohamedveron/go-microservice-template/persistence"

type Service struct {
	persistence *persistence.Persistence
}

func NewService(persistence *persistence.Persistence) *Service {

	return &Service{
		persistence: persistence,
	}
}
