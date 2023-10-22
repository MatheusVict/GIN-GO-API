package service

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/test/mocks"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

func TestUserDomainService_LoginUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository: repository}

	t.Run("when_calling_repository_returns_errors", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 5)
		userDomain.SetID(id)

		userDomainMock := model.NewUserDomain(
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetName(), userDomain.GetAge(),
		)
		userDomainMock.EncryptPassword()

		repository.EXPECT().FindUserByEmailAndPassword(userDomain.GetEmail(), userDomainMock.GetPassword()).Return(
			nil,
			errorsHandle.NewInternalServerError("error trying to find user by email and password"),
		)

		user, token, err := service.LoginUserService(userDomain)

		assert.Nil(mt, user)
		assert.Empty(mt, token)
		assert.NotNil(mt, err)
		assert.EqualValues(mt, err.Message, "error trying to find user by email and password")
	})

	t.Run("when_calling_create_token_returns_error", func(mt *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(ctrl)
		userDomainMock.EXPECT().GetEmail().Return("test@tes.com")
		userDomainMock.EXPECT().GetPassword().Return("test")
		userDomainMock.EXPECT().EncryptPassword()

		userDomainMock.EXPECT().GenerateToken().Return(
			"",
			errorsHandle.NewInternalServerError("error trying to create token"),
		)
		repository.EXPECT().FindUserByEmailAndPassword(
			"test@tes.com",
			"test",
		).Return(userDomainMock, nil)

		user, token, err := service.LoginUserService(userDomainMock)

		assert.Nil(mt, user)
		assert.Empty(mt, token)
		assert.NotNil(mt, err)
		assert.EqualValues(mt, err.Message, "error trying to create token")

	})

	t.Run("when_email_and_password_is_valid_return_success", func(mt *testing.T) {
		id := primitive.NewObjectID().Hex()
		secret := "test"
		os.Setenv("JWT_SECRET_KEY", secret)
		defer os.Clearenv()
		userDomain := model.NewUserDomain("test@test.com", "test", "test", 5)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(), gomock.Any()).Return(
			userDomain, nil,
		)

		user, token, err := service.LoginUserService(userDomain)

		assert.Nil(mt, err)
		assert.EqualValues(mt, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(mt, user.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(mt, user.GetName(), userDomain.GetName())

		tokenReturned, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}

			return nil, errorsHandle.NewBadRequestError("invalid token")
		})

		_, ok := tokenReturned.Claims.(jwt.MapClaims)
		if !ok || !tokenReturned.Valid {
			t.FailNow()
			return
		}
	})
}
