package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(userDomain model.UserDomainInterface) *errorsHandle.ErrorsHandle
	UpdateUser(string, model.UserDomainInterface) *errorsHandle.ErrorsHandle
	FindUser(string) (*model.UserDomainInterface, *errorsHandle.ErrorsHandle)
	DeleteUser(string) *errorsHandle.ErrorsHandle
}
