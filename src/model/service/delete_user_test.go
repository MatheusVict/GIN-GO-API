package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		repository.EXPECT().DeleteUser(id).Return(nil)

		err := service.DeleteUser(id)

		assert.Nil(t, err)
	})

	t.Run("when_sending_an_invalid_userID_returns_error", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(
			errorsHandle.NewInternalServerError("error trying to delete user"),
		)

		err := service.DeleteUser(id)

		assert.NotNil(t, err)
		assert.EqualValues(mt, err.Message, "error trying to delete user")
	})
}
