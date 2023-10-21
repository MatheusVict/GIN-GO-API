package model

import "github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	SetID(id string)
	GetID() string

	EncryptPassword()
	GetJSONValue() (string, error)
	GenerateToken() (string, *errorsHandle.ErrorsHandle)
}

func NewUserDomain(
	email,
	password,
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}
func NewUserLoginDomain(
	email string,
	password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
