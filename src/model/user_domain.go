package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	EncryptPassword()
}

func NewUserDomain(
	email,
	password,
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email,
		password,
		name,
		age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	Age      int8
}

func (user *userDomain) GetEmail() string {
	return user.email
}
func (user *userDomain) GetPassword() string {
	return user.password
}
func (user *userDomain) GetName() string {
	return user.name
}
func (user *userDomain) GetAge() int8 {
	return user.Age
}

func (user *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.password))
	user.password = hex.EncodeToString(hash.Sum(nil))
}
