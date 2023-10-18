package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
)

func NewUserDomain(
	email,
	password,
	name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email,
		password,
		name,
		age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (user *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *errorsHandle.ErrorsHandle
	UpdateUser(string) *errorsHandle.ErrorsHandle
	FindUser(string) (*UserDomain, *errorsHandle.ErrorsHandle)
	DeleteUser(string) *errorsHandle.ErrorsHandle
}
