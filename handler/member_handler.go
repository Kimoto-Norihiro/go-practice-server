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
	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = mh.useCase.CreateMember(c, m)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (mh *MemberHandler) ShowMember(c *gin.Context) {
	id := c.Param("id")
	m, err := mh.useCase.ShowMember(c, id)
	if err != nil {
		c.JSON(500, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data":  m,
		"error": nil,
	})
}

func (mh *MemberHandler) DeleteMember(c *gin.Context) {
	id := c.Param("id")
	err := mh.useCase.DeleteMember(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (mh *MemberHandler) UpdateMember(c *gin.Context) {
	id := c.Param("id")
	var m model.Member
	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = mh.useCase.UpdateMember(c, id, m)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}


