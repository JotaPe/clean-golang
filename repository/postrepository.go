package repository

import (
	"clean-arch/entity"
)

// Post repo interface for clean setup
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() (*[]entity.Post, error)
}
