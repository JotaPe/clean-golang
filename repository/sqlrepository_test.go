package repository

import (
	"clean-arch/entity"

	"database/sql"
	"testing"

	"github.com/gofrs/uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Mock struct
type GormSuite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

/*
func TestGenerateEveryRun(t *testing.T) {
	suite := &GormSuite{}
	var (
		db  *sql.DB
		err error
	)

	db, suite.mock, err = sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	if db == nil {
		t.Error("Mock db is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	suite.db, _ = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	suite.mock.MatchExpectationsInOrder(false)
	suite.mock.ExpectBegin()
	t.Log(suite.db)

	suite.mock.ExpectCommit()
	suite.db.AutoMigrate(&Post{})

	defer db.Close()

	postRepository := NewSqlRepository(suite.db)

	generateId, _ := uuid.NewV4()
	suite.mock.ExpectCommit()
	generate, err := postRepository.Save(&entity.Post{ID: generateId.String(), Title: "Test Title", Text: "Test Text"})
	if err != nil {
		t.Fatal(err)
	}
	suite.mock.ExpectationsWereMet()

	t.Logf("Generated post: %v", generate)
}

func TestFindAll(t *testing.T) {
	suite := &GormSuite{}
	var (
		db  *sql.DB
		err error
	)

	db, suite.mock, err = sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	if db == nil {
		t.Error("Mock db is null")
	}
	if suite.mock == nil {
		t.Error("sqlmock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	suite.db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	t.Log(suite.db)

	suite.db.AutoMigrate(&Post{})

	defer db.Close()

	postRepository := NewSqlRepository(suite.db)
	if err != nil {
		t.Error(err)
	}

	generateID, _ := uuid.NewV4()
	generate, err := postRepository.Save(&entity.Post{ID: generateID.String(), Title: "Test Title", Text: "Test Text"})
	t.Logf("Generated for FINDALL: %v", generate)

	posts, err := postRepository.FindAll()
	if err != nil {
		t.Error(err)
	}
	for i, s := range *posts {
		t.Log(i, s)
	}
}
*/
