package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/go-practice-server/model"
)

type UseCase interface {
	CreateMember(c *gin.Context, m model.Member) error
}