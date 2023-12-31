package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"log"
)

func (u *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("CreateUser")
	user, err := u.FindUserByEmailService(userDomain.GetEmail())
	if user != nil {
		return nil, errorsHandle.NewBadRequestError("Email is already registered in another account")
	}
	if err != nil {
		log.Println("Error on search user")
	}
	userDomain.EncryptPassword()
	log.Println(userDomain.GetPassword())

	userDomainRepository, err := u.repository.CreateUser(userDomain)
	if err != nil {
		log.Println("Error on createUser mode: ", err)
		return nil, err
	}
	return userDomainRepository, nil
}
