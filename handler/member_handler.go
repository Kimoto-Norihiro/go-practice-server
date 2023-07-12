package handler

import (
	"github.com/Kimoto-Norihiro/go-practice-server/model"
	"github.com/Kimoto-Norihiro/go-practice-server/usecase"
	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	useCase usecase.UseCase
}

func NewMemberHandler(u usecase.UseCase) *MemberHandler {
	return &MemberHandler{u}
}

func (mh *MemberHandler) CreateMember(c *gin.Context) {
	var m model.Member
	c.BindJSON(&m)
	mh.useCase.CreateMember(c, m)
}


