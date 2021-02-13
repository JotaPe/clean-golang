package repository

import (
	"clean-arch/entity"
	"clean-arch/repository"
	"testing"
)

var (
	postRepository repository.PostRepository = repository.NewSqlRepository()
)

func TestFindAll(t *testing.T) {
	posts, err := postRepository.FindAll()
	if err != nil {
		t.Error(err)
	}
	for i, s := range *posts {
		t.Log(i, s)
	}
}

func TestGenerateEveryRun(t *testing.T) {
	generate, err := postRepository.Save(&entity.Post{ID: "1234132451", Title: "Test Title", Text: "Test Text"})
	if err != nil {
		t.Error(err)
	}
	t.Log(generate)
}
