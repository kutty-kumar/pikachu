package svc

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type UserRepositoryMock struct {
	mock.Mock

}


func TestUserCreate(t *testing.T){
	testUserRepo := new(UserRepositoryMock)
	testUserRepo.On("Create", mock.Anything).Return()
}