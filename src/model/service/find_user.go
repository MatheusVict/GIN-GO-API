package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"log"
)

func (user *userDomainService) FindUserByIDService(
	id string,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("Init on FindUserByIDService")
	return user.repository.FindUserByID(id)
}

func (user *userDomainService) FindUserByEmailService(
	email string,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("Init on FindUserByEmailService")

	return user.repository.FindUserByEmail(email)
}
func (user *userDomainService) findUserByEmailAndPasswordService(
	email string,
	password string,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("Init on FindUserByEmailService")

	return user.repository.FindUserByEmailAndPassword(email, password)
}
