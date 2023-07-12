package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/go-practice-server/db"
	"github.com/Kimoto-Norihiro/go-practice-server/repository"
	"github.com/Kimoto-Norihiro/go-practice-server/usecase"
	"github.com/Kimoto-Norihiro/go-practice-server/handler"
)

func main() {
	r := gin.Default()

	const dns = "n000r111:password@tcp(localhost:3306)/go_practice_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := database.NewMySql(dns)
	if err != nil {
		panic(err)
	}
	repo := repository.NewMemberRepository(db)
	usecase := usecase.NewMemberUseCase(repo)
	handler := handler.NewMemberHandler(usecase)

	r.POST("/member", handler.CreateMember)
	r.GET("/member/:id", handler.ShowMember)
	r.DELETE("/member/:id", handler.DeleteMember)
	r.PUT("/member/:id", handler.UpdateMember)

	r.Run()
}