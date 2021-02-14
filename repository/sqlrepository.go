package repository

import (
	"clean-arch/entity"

	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repo struct{}

var DB *gorm.DB

type Post struct {
	gorm.Model
	ID        string
	Title     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSqlRepository() PostRepository {
	var err error
	dsn := "host=localhost user=postgres password=docker dbname=pqgotest port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection with SQL database open")

	DB.AutoMigrate(&Post{})
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
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
	result := DB.Create(&entityInsert)
	if result.Error != nil {
		log.Fatal(err)
	}

	return &entity, nil
}

func (*repo) FindAll() (*[]entity.Post, error) {
	var entityToFind []Post
	var entities []entity.Post
	result := DB.Find(&entityToFind).Scan(&entities)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return &entities, nil
}
