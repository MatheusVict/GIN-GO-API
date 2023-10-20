package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	EncryptPassword()
	GetJSONValue() (string, error)

	SetID(id string)
}

func NewUserDomain(
	email,
	password,
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	ID       string
	Email    string `json:"email"`
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) SetID(id string) {

}

func (ud *userDomain) GetJSONValue() (string, error) {
	sliceOfBytes, err := json.Marshal(ud)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(sliceOfBytes), nil
}

func (user *userDomain) GetEmail() string {
	return user.Email
}
func (user *userDomain) GetPassword() string {
	return user.Password
}
func (user *userDomain) GetName() string {
	return user.Name
}
func (user *userDomain) GetAge() int8 {
	return user.Age
}

func (user *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))
}
