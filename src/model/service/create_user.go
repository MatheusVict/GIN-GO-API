package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"log"
)

func (u *userDomainService) CreateUser(userDomain model.UserDomainInterface) *errorsHandle.ErrorsHandle {
	log.Println("CreateUser")
	userDomain.EncryptPassword()
	log.Println(userDomain.GetPassword())
	return nil
}
