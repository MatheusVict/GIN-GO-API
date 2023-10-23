package controller

import (
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUserControllerInterface_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("when_email_is_invalid_returns_error", func(mt *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{Key: "userEmail", Value: "TEST_ERROR"},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(mt, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when_email_is_valid_returns_success", func(mt *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{Key: "userEmail", Value: "test@test.com"},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").Return(
			model.NewUserDomain("test@test.com", "hello", "mat", 8), nil,
		)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
func TestUserControllerInterface_FindUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("when_id_is_invalid_returns_error", func(mt *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{Key: "userId", Value: "test"},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(mt, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when_id_is_valid_returns_success", func(mt *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{Key: "userId", Value: id},
		}

		service.EXPECT().FindUserByIDService(id).Return(
			model.NewUserDomain("test@test.com", "hello", "mat", 8), nil,
		)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.URL.RawQuery = u.Encode()
}
