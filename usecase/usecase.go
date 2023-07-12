package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/go-practice-server/model"
)

type UseCase interface {
	CreateMember(c *gin.Context, m model.Member) error
	ShowMember(c *gin.Context, id string) (model.Member, error)
	DeleteMember(c *gin.Context, id string) error
	UpdateMember(c *gin.Context, id string, m model.Member) error
}