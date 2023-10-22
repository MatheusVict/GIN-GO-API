package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_CreateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_returns_error", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 5)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUserService(userDomain)

		assert.Nil(mt, user)
		assert.NotNil(mt, err)
		assert.EqualValues(mt, err.Message, "Email is already registered in another account")
	})

	t.Run("when_user_is_not_registered_returns_error", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 5)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			nil,
			nil,
		)

		repository.EXPECT().CreateUser(userDomain).Return(
			nil,
			errorsHandle.NewInternalServerError("error trying to create user"),
		)

		user, err := service.CreateUserService(userDomain)

		assert.Nil(mt, user)
		assert.NotNil(mt, err)
		assert.EqualValues(mt, err.Message, "error trying to create user")
	})

	t.Run("when_user_is_not_registered_returns_success", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 5)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			nil,
			nil,
		)

		repository.EXPECT().CreateUser(userDomain).Return(
			userDomain,
			nil,
		)

		user, err := service.CreateUserService(userDomain)

		assert.NotNil(mt, user)
		assert.Nil(mt, err)
		assert.EqualValues(mt, userDomain.GetName(), user.GetName())
		assert.EqualValues(mt, userDomain.GetPassword(), user.GetPassword())
	})

}
