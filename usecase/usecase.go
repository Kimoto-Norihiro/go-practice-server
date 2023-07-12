package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/go-practice-server/model"
)

type UseCase interface {
	CreateMember(c *gin.Context, m model.Member) error
	ShowMember(c *gin.Context, id uint) (model.Member, error)
	DeleteMember(c *gin.Context, id uint) error
	UpdateMember(c *gin.Context, id uint, m model.Member) error
}
