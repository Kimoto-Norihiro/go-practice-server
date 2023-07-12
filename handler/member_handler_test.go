package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/go-practice-server/handler"
	"github.com/Kimoto-Norihiro/go-practice-server/mock"
	"github.com/Kimoto-Norihiro/go-practice-server/model"
)

func setupTestContext(t *testing.T, method, url string, body interface{}) (*gin.Context, *gin.Engine, *gomock.Controller, *mock.MockUseCase) {
	ginCtx, r := gin.CreateTestContext(httptest.NewRecorder())

	ctrl := gomock.NewController(t)

	usecase := mock.NewMockUseCase(ctrl)

	ginCtx.Request = httptest.NewRequest(method, url, nil)
	if body != nil {
		bodyBytes, _ := json.Marshal(body)
		ginCtx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	}

	return ginCtx, r, ctrl, usecase
}

func TestHandlerCreateMember(t *testing.T) {
	expectedMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}

	ginCtx, _, ctrl, mu := setupTestContext(t, "POST", "/member", expectedMember)
	defer ctrl.Finish()

	mu.EXPECT().CreateMember(ginCtx, expectedMember).Return(nil)

	mh := handler.NewMemberHandler(mu)

	mh.CreateMember(ginCtx)
	
	if ginCtx.Writer.Status() != http.StatusOK {
		t.Errorf("status code should be 200, but got: %v", ginCtx.Writer.Status())
	}
}

func TestHandlerShowMember(t *testing.T) {
	expectedMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}

	ginCtx, _, ctrl, mu := setupTestContext(t, "GET", "/member/1", nil)
	ginCtx.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	defer ctrl.Finish()

	mu.EXPECT().ShowMember(ginCtx, uint(1)).Return(expectedMember, nil)

	mh := handler.NewMemberHandler(mu)

	mh.ShowMember(ginCtx)

	if ginCtx.Writer.Status() != http.StatusOK {
		t.Errorf("status code should be 200, but got: %v", ginCtx.Writer.Status())
	}
}

func TestHandlerUpdateMember(t *testing.T) {
	expectedMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}

	ginCtx, _, ctrl, mu := setupTestContext(t, "PUT", "/member/1", expectedMember)
	ginCtx.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	defer ctrl.Finish()

	mu.EXPECT().UpdateMember(ginCtx, uint(1), expectedMember).Return(nil)

	mh := handler.NewMemberHandler(mu)

	mh.UpdateMember(ginCtx)

	if ginCtx.Writer.Status() != http.StatusOK {
		t.Errorf("status code should be 200, but got: %v", ginCtx.Writer.Status())
	}
}

func TestHandlerDeleteMember(t *testing.T) {
	ginCtx, _, ctrl, mu := setupTestContext(t, "DELETE", "/member/1", nil)
	ginCtx.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	defer ctrl.Finish()

	mu.EXPECT().DeleteMember(ginCtx, uint(1)).Return(nil)

	mh := handler.NewMemberHandler(mu)

	mh.DeleteMember(ginCtx)

	if ginCtx.Writer.Status() != http.StatusOK {
		t.Errorf("status code should be 200, but got: %v", ginCtx.Writer.Status())
	}
}
