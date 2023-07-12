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

	const dns = "hello"
	db, err := database.NewMySql(dns)
	if err != nil {
		panic(err)
	}
	repo := repository.NewMemberRepository(db)
	usecase := usecase.NewMemberUseCase(repo)
	handler := handler.NewMemberHandler(usecase)

	r.POST("/member", handler.CreateMember)

	r.Run()
}