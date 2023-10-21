package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"log"
)

func (u *userDomainService) LoginUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("LoginUser")
	userDomain.EncryptPassword()

	user, err := u.findUserByEmailAndPasswordService(userDomain.GetEmail(), userDomain.GetPassword())

	if err != nil {
		return nil, err
	}

	return user, nil
}
