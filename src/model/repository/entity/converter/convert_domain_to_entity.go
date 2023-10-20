package converter

import (
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
		Password: domain.GetPassword(),
	}

}
