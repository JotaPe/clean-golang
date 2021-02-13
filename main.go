package main

import (
	"clean-arch/repository"
	"testing"
)

var (
	postRepository repository.PostRepository = repository.NewSqlRepository()
)

func main() {
}
