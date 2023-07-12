package usecase

import (
	"github.com/Kimoto-Norihiro/go-practice-server/model"
	"github.com/Kimoto-Norihiro/go-practice-server/repository"
	"github.com/gin-gonic/gin"
)

type MemberUseCase struct {
	repository repository.Repository
}

func NewMemberUseCase(r repository.Repository) *MemberUseCase {
	return &MemberUseCase{r}
}

func (mu *MemberUseCase) CreateMember(c *gin.Context, m model.Member) error {
	err := mu.repository.CreateMember(m)
	if err != nil {
		return err
	}
	return nil
}

func (mu *MemberUseCase) ShowMember(c *gin.Context, id string) (model.Member, error) {
	m, err := mu.repository.ShowMember(id)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (mu *MemberUseCase) DeleteMember(c *gin.Context, id string) error {
	err := mu.repository.DeleteMember(id)
	if err != nil {
		return err
	}
	return nil
}

func (mu *MemberUseCase) UpdateMember(c *gin.Context, id string, m model.Member) error {
	err := mu.repository.UpdateMember(id, m)
	if err != nil {
		return err
	}
	return nil
}