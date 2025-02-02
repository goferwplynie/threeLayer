package main

import (
	"fmt"
	"projekt/internal/repository"
)

func main() {
	repo := repository.New()

	user := repo.Users[0]

	fmt.Println(user)
}
