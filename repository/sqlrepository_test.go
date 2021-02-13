package repository

import (
	"clean-arch/entity"

	"testing"
)

func TestSqlSave(t *testing.T) {
	postRepository := repository.NewSqlRepository()
	post := entity.Post{
		ID:    "123341",
		Title: "Test Post",
		Text:  "Test Post Text",
	}

	postSaved := postRepository.Save(post)

	t.Log(postSaved)
}
