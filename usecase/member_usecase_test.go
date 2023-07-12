package usecase_test

import (
	"testing"
	"net/http/httptest"
	"bytes"
	"encoding/json"
	"io/ioutil"

	"gorm.io/gorm"
	"github.com/golang/mock/gomock"
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/go-practice-server/mock"
	"github.com/Kimoto-Norihiro/go-practice-server/usecase"
	"github.com/Kimoto-Norihiro/go-practice-server/model"
)

func setupTestContext(t *testing.T, method, url string, body interface{}) (*gin.Context, *gomock.Controller, *mock.MockRepository) {
	ginCtx, _ := gin.CreateTestContext(httptest.NewRecorder())

	ctrl := gomock.NewController(t)

	repo := mock.NewMockRepository(ctrl)

	ginCtx.Request = httptest.NewRequest(method, url, nil)
	if body != nil {
		bodyBytes, _ := json.Marshal(body)
		ginCtx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
	}

	return ginCtx, ctrl, repo
}

func TestUseCaseCreateUser(t *testing.T) {
	expectedMember := model.Member{
		Name: "test",
	}

	ginCtx, ctrl, repo := setupTestContext(t, "POST", "/member", expectedMember)
	defer ctrl.Finish()

	repo.EXPECT().CreateMember(expectedMember).Return(nil)

	mu := usecase.NewMemberUseCase(repo)

	err := mu.CreateMember(ginCtx, expectedMember)
	if err != nil {
		t.Errorf("error should be nil, but got: %v", err)
	}
}

func TestUseCaseShowUser(t *testing.T) {
	expectedMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}

	ginCtx, ctrl, repo := setupTestContext(t, "GET", "/member/1", nil)
	defer ctrl.Finish()

	repo.EXPECT().ShowMember(expectedMember.ID).Return(expectedMember, nil)

	mu := usecase.NewMemberUseCase(repo)

	_, err := mu.ShowMember(ginCtx, expectedMember.ID)
	if err != nil {
		t.Errorf("error should be nil, but got: %v", err)
	}
}

func TestUseCaseDeleteUser(t *testing.T) {
	expectedMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}

	ginCtx, ctrl, repo := setupTestContext(t, "DELETE", "/member/1", nil)
	defer ctrl.Finish()
	
	repo.EXPECT().DeleteMember(expectedMember.ID).Return(nil)

	mu := usecase.NewMemberUseCase(repo)

	err := mu.DeleteMember(ginCtx, expectedMember.ID)
	if err != nil {
		t.Errorf("error should be nil, but got: %v", err)
	}
}

func TestUseCaseUpdateUser(t *testing.T) {
	expectedMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}

	ginCtx, ctrl, repo := setupTestContext(t, "PUT", "/member/1", expectedMember)
	defer ctrl.Finish()
	
	repo.EXPECT().UpdateMember(expectedMember.ID, expectedMember).Return(nil)

	mu := usecase.NewMemberUseCase(repo)

	err := mu.UpdateMember(ginCtx, expectedMember.ID, expectedMember)
	if err != nil {
		t.Errorf("error should be nil, but got: %v", err)
	}
}