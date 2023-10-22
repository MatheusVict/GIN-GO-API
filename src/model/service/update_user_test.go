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

func TestUserDomainService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 5)
		userDomain.SetID(id)
		repository.EXPECT().UpdateUser(id, userDomain).Return(nil)

		err := service.UpdateUser(id, userDomain)

		assert.Nil(t, err)
	})
	t.Run("when_sending_an_invalid_user_and_userID_returns_error", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 5)
		userDomain.SetID(id)
		repository.EXPECT().UpdateUser(id, userDomain).Return(
			errorsHandle.NewInternalServerError("error trying to update user"),
		)

		err := service.UpdateUser(id, userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(mt, err.Message, "error trying to update user")
	})

}
