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

func TestUserDomainService_FindUserByIDService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 5)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByID(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByIDService(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		repository.EXPECT().FindUserByID(id).Return(nil, errorsHandle.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByIDService(id)
		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")

	})
}

func TestUserDomainService_FindUserByEmailService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"
		userDomain := model.NewUserDomain(email, "test", "test", 5)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmailService(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(mt *testing.T) {
		email := "test@test.com"
		repository.EXPECT().FindUserByEmail(email).Return(
			nil,
			errorsHandle.NewNotFoundError("user not found"),
		)

		userDomainReturn, err := service.FindUserByEmailService(email)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")

	})
}

func TestUserDomainService_FindUserByEmailAndPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository: repository}

	t.Run("when_exists_an_user_returns_success", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"
		password := "test"
		userDomain := model.NewUserDomain(email, password, "test", 5)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		userDomainReturn, err := service.findUserByEmailAndPasswordService(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(mt *testing.T) {
		email := "test@test.com"
		password := "test"
		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(
			nil,
			errorsHandle.NewNotFoundError("user not found"),
		)

		userDomainReturn, err := service.findUserByEmailAndPasswordService(email, password)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")

	})
}
