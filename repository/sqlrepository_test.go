package repository

import (
	"clean-arch/entity"
	"clean-arch/repository"
	"github.com/google/uuid"
	"testing"
)

var (
	postRepository repository.PostRepository = repository.NewSqlRepository()
)

func TestGenerateEveryRun(t *testing.T) {
	generateId := uuid.NewString()
	generate, err := postRepository.Save(&entity.Post{ID: generateId, Title: "Test Title", Text: "Test Text"})
	if err != nil {
		t.Error(err)
	}
	t.Log(generate)
}

func TestFindAll(t *testing.T) {
	posts, err := postRepository.FindAll()
	if err != nil {
		t.Error(err)
	}
	for i, s := range *posts {
		t.Log(i, s)
	}
}
