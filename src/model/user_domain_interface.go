package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	SetID(id string)
	GetID() string

	EncryptPassword()
	GetJSONValue() (string, error)
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
