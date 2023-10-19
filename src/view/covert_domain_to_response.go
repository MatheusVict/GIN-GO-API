package view

import (
	"github.com/MatheusVict/User-Register-GO/src/controller/model/response"
	"github.com/MatheusVict/User-Register-GO/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Age:   userDomain.GetAge(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
	}
}
