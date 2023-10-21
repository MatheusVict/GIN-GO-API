package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"log"
)

func (user *userDomainService) DeleteUser(userId string) *errorsHandle.ErrorsHandle {
	log.Println("Delete")

	err := user.repository.DeleteUser(userId)
	if err != nil {
		log.Println("Error on delete user: ", err)
		return err
	}

	log.Println("User deleted successfully")
	return nil
}
