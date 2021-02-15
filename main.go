package main

import (
	"clean-arch/entity"
	"clean-arch/repository"
	"github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=docker dbname=pqgotest port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection with SQL database open")

	DB.AutoMigrate(&repository.Post{})

	postRepository := repository.NewSqlRepository(DB)
	generateID, _ := uuid.NewV4()
	generate, _ := postRepository.Save(&entity.Post{ID: generateID.String(), Title: "Test Title", Text: "Test Text"})
	log.Printf("Generated for FINDALL: %v", generate)

	posts, err := postRepository.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, s := range *posts {
		log.Println(i, s)
	}
}
