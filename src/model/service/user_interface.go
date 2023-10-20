package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(userDomain model.UserDomainInterface) *errorsHandle.ErrorsHandle
	UpdateUser(string, model.UserDomainInterface) *errorsHandle.ErrorsHandle
	FindUser(string) (*model.UserDomainInterface, *errorsHandle.ErrorsHandle)
	DeleteUser(string) *errorsHandle.ErrorsHandle
}
