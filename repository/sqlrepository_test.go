package repository

import (
	"clean-arch/entity"

	//"database/sql"
	"testing"

	"github.com/gofrs/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(t *testing.T) {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateEveryRun(t *testing.T) {
	postRepository := NewSqlRepository(db)

	generateId, _ := uuid.NewV4()
	generate, err := postRepository.Save(&entity.Post{ID: generateId.String(), Title: "Test Title", Text: "Test Text"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Generated post: %v", generate)
}

func TestFindAll(t *testing.T) {
	postRepository := NewSqlRepository(db)

	posts, err := postRepository.FindAll()
	if err != nil {
		t.Error(err)
	}
	for i, s := range *posts {
		t.Log(i, s)
	}
}
