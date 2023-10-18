package model

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"log"
)

func (u *UserDomain) CreateUser() *errorsHandle.ErrorsHandle {
	log.Println("CreateUser")
	u.EncryptPassword()
	log.Println()
	return nil
}
