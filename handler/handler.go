package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateMember(c *gin.Context)
	ShowMember(c *gin.Context)
	DeleteMember(c *gin.Context)
	UpdateMember(c *gin.Context)
}