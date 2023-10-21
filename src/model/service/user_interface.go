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
	CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *errorsHandle.ErrorsHandle)
	UpdateUser(string, model.UserDomainInterface) *errorsHandle.ErrorsHandle
	FindUserByEmailService(email string) (model.UserDomainInterface, *errorsHandle.ErrorsHandle)
	FindUserByIDService(id string) (model.UserDomainInterface, *errorsHandle.ErrorsHandle)
	DeleteUser(string) *errorsHandle.ErrorsHandle
	LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *errorsHandle.ErrorsHandle)
}
