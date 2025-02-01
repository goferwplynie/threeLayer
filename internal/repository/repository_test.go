package repository

import (
	"projekt/internal/models"
	"testing"
)

func TestCreateRepo(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Error("repo not created")
	}
}

func TestAddUser(t *testing.T) {
	var exampleUser = models.User{
		FirstName: "Dawid",
		LastName:  "Markiewicz",
		Age:       15,
		Group:     "admin",
	}
	repo := New()
	id := repo.NextId
	exampleUser.Id = id
	repo.AddUser(exampleUser)

	if repo.Users[id] != exampleUser {
		t.Errorf("user is not on next id\n%v", repo.Users)
	}
}
