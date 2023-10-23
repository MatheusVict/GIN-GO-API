package controller

import (
	"encoding/json"
	"github.com/MatheusVict/User-Register-GO/src/controller/model/request"
	"github.com/MatheusVict/User-Register-GO/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserControllerInterface_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("when_validation_got_an_error", func(mt *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{}
		userRequest := request.UserRequest{
			Email:    "EMAIL_ERROR",
			Password: "test",
			Name:     "hello",
			Age:      9,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, param, url.Values{}, "POST", stringReader)
		controller.DeleteUser(context)

		assert.EqualValues(mt, http.StatusBadRequest, recorder.Code)
	})
}
