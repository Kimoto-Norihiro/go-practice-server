package repository_test

import (
	"testing"
	"fmt"

	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/go-practice-server/db"
	"github.com/Kimoto-Norihiro/go-practice-server/model"
	"github.com/Kimoto-Norihiro/go-practice-server/repository"
)

func setupTest() (*repository.MemberRepository, *gorm.DB, func()) {
	// Create a test database connection
	const dns = "n000r111:password@tcp(localhost:3306)/go_practice_server_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := database.NewMySql(dns)
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.Member{})
	if err != nil {
		panic(fmt.Errorf("failed to migrate database: %v", err))
	}

	// Create the repository
	repo := repository.NewMemberRepository(db)

	// Define a cleanup function to drop the member table after the tests
	cleanup := func() {
		db.Migrator().DropTable(&model.Member{})
		sqlDB, err := db.DB()
		if err != nil {
			panic(fmt.Errorf("failed to close database: %v", err))
		}
		sqlDB.Close()
	}

	return repo, db, cleanup
}

func TestRepositoryCreateUser(t *testing.T) {
	var err error
	repo, db, cleanup := setupTest()
	defer cleanup()
	// expect member
	expectMember := model.Member{
		Name: "test",
	}
	// create test user
	err = repo.CreateMember(expectMember)
	if err != nil {
		t.Errorf("fail to create user: %v", err)
	}
	// get test user
	var result model.Member
	db.First(&result, "name = ?", "test")

	// check creation
	if result.Name != expectMember.Name {
		t.Errorf("fail to get user: %v", err)
	}
}

func TestRepositoryShowUser(t *testing.T) {
	var err error
	repo, _, cleanup := setupTest()
	defer cleanup()
	// expect member
	expectMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}
	// create test user
	err = repo.CreateMember(expectMember)
	if err != nil {
		t.Errorf("fail to create user: %v", err)
	}
	// get test user
	result, err := repo.ShowMember(expectMember.ID)
	if err != nil {
		t.Errorf("fail to get user: %v", err)
	}
	// check show
	if result.Name != expectMember.Name {
		t.Errorf("fail to get user: %v", err)
	}
}

func TestRepositoryDeleteUser(t *testing.T) {
	repo, _, cleanup := setupTest()
	defer cleanup()
	// expect member
	expectMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}
	// create test user
	err := repo.CreateMember(expectMember)
	if err != nil {
		t.Errorf("fail to create user: %v", err)
	}
	// delete test user
	err = repo.DeleteMember(expectMember.ID)
	if err != nil {	
		t.Errorf("fail to delete user: %v", err)
	}
}

func TestRepositoryUpdateUser(t *testing.T) {
	repo, _, cleanup := setupTest()
	defer cleanup()
	// before update member
	beforeMember := model.Member{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}
	// after update member
	afterMember := model.Member{
		Name: "update",
	}
	// create test user
	err := repo.CreateMember(beforeMember)
	if err != nil {
		t.Errorf("fail to create user: %v", err)
	}
	// update test user
	err = repo.UpdateMember(beforeMember.ID, afterMember)
	if err != nil {
		t.Errorf("fail to update user: %v", err)
	}
	// get test user
	result, err := repo.ShowMember(beforeMember.ID)
	if err != nil {
		t.Errorf("fail to get user: %v", err)
	}
	// check update
	if result.Name != afterMember.Name {
		t.Errorf("fail to get user: %v", err)
	}
}