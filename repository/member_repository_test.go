package repository_test

import (
	"testing"

	"github.com/Kimoto-Norihiro/go-practice-server/db"
	"github.com/Kimoto-Norihiro/go-practice-server/model"
	"github.com/Kimoto-Norihiro/go-practice-server/repository"
)

func TestRepositoryCreateUser(t *testing.T) {
	// テスト用のDBを作成
	const dns = "n000r111:password@tcp(localhost:3306)/go_practice_server_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := database.NewMySql(dns)
	if err != nil {
		t.Errorf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&model.Member{})
	if err != nil {
		t.Errorf("failed to migrate database: %v", err)
	}

	// テスト用のリポジトリを作成
	repo := repository.NewMemberRepository(db)

	// expect member
	expectMember := model.Member{
		Name: "test",
	}

	// テスト用のユーザーを作成
	err = repo.CreateMember(expectMember)
	if err != nil {
		t.Errorf("fail to create user: %v", err)
	}

	// テスト用のユーザーを取得
	var result model.Member
	db.First(&result, "name = ?", "test")

	// check creation
	if result.Name != expectMember.Name {
		t.Errorf("fail to get user: %v", err)
	}

	// delete member table
	db.Migrator().DropTable(&model.Member{})
}