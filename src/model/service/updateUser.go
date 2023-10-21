package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"log"
)

func (user *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *errorsHandle.ErrorsHandle {
	log.Println("UpdateUser")

	err := user.repository.UpdateUser(userId, userDomain)
	if err != nil {
		log.Println("Error on createUser mode: ", err)
		return err
	}

	log.Println("User updated successfully")
	return nil
}
