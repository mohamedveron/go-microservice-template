package test

import (
	"github.com/TheSocialLoanCompany/go-microservice-template/persistence"
	"github.com/TheSocialLoanCompany/go-microservice-template/service"
	"testing"
)

func TestGetAccounts(t *testing.T){

	persistenceLayer := persistence.NewPersistence(nil)
	serviceLayer := service.NewService(persistenceLayer)

	_, err := serviceLayer.GetAccounts()

	if err != nil{
		t.Errorf("Failed to get accounts")
	}

}