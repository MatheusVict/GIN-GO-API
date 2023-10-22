package repository

import (
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DATABASE", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_valid_user_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"email@mail.com",
			"test",
			"test",
			8,
		)
		domain.SetID(primitive.NewObjectID().Hex())
		err := repo.UpdateUser(domain.GetID(), domain)

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database_when_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"email@mail.com",
			"test",
			"test",
			8,
		)
		domain.SetID(primitive.NewObjectID().Hex())
		err := repo.UpdateUser(domain.GetID(), domain)
		assert.NotNil(t, err)
	})
}
