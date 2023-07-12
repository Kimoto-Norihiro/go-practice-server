package usecase

import (
	"github.com/Kimoto-Norihiro/go-practice-server/model"
	"github.com/Kimoto-Norihiro/go-practice-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MemberUseCase struct {
	repository repository.Repository
	validate   *validator.Validate
}

func NewMemberUseCase(r repository.Repository) *MemberUseCase {
	return &MemberUseCase{
		repository: r,
		validate:   validator.New(),
	}
}

func (mu *MemberUseCase) CreateMember(c *gin.Context, m model.Member) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.CreateMember(m)
}

func (mu *MemberUseCase) ShowMember(c *gin.Context, id uint) (model.Member, error) {
	return mu.repository.ShowMember(id)
}

func (mu *MemberUseCase) DeleteMember(c *gin.Context, id uint) error {
	return mu.repository.DeleteMember(id)
}

func (mu *MemberUseCase) UpdateMember(c *gin.Context, id uint, m model.Member) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.UpdateMember(id, m)
}
