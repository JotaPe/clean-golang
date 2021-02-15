package repository

import (
	"clean-arch/entity"

	"log"
	"time"

	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

// Post custom struct for SQL transactions
type Post struct {
	gorm.Model
	entity.Post
	ID        string
	Title     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewSqlRepository is used to generate the common interface for PostRepository
func NewSqlRepository(db *gorm.DB) PostRepository {
	return &repo{DB: db}
}

func (repo *repo) Save(post *entity.Post) (*entity.Post, error) {
	var err error
	entity := entity.Post{
		ID:    post.ID,
		Title: post.Title,
		Text:  post.Text,
	}
	entityInsert := Post{
		ID:    post.ID,
		Title: post.Title,
		Text:  post.Text,
	}
	result := repo.DB.Create(&entityInsert)
	if result.Error != nil {
		log.Fatal(err)
	}

	return &entity, nil
}

func (repo *repo) FindAll() (*[]entity.Post, error) {
	var entityToFind []Post
	var entities []entity.Post
	result := repo.DB.Find(&entityToFind).Scan(&entities)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return &entities, nil
}
