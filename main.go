package main

import (
	"clean-arch/repository"
)

var (
	postRepository repository.PostRepository = repository.NewSqlRepository()
)

func main() {
}
