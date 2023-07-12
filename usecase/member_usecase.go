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