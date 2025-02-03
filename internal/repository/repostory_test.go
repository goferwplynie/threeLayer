package repository

import (
	"projekt/internal/models"
	"testing"
)

func CreateRepo() *Repository {
	user := models.User{
		Id:        1,
		FirstName: "Dawid",
		LastName:  "Markiewicz",
		Age:       15,
		Group:     "admin",
	}

	return &Repository{
		Users: map[int]models.User{
			1: user,
		},
		NextId: 2,
	}
}

func TestGetUser(t *testing.T) {
	repo := CreateRepo()

	user, err := repo.GetUserById(1)

	if err != nil {
		t.Errorf("error thrown: %v", err)
	}

	if user.FirstName != "Dawid" {
		t.Errorf("wrong user returned: %v", user)
	}
}

func TestGetAllUsers(t *testing.T) {
	repo := CreateRepo()

	users := repo.GetAll()

	if users[0].FirstName != "Dawid" {
		t.Errorf("wrong user returned: %v", users)
	}
}

func TestAddUser(t *testing.T) {
	repo := CreateRepo()

	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       24,
		Group:     "premium",
	}

	repo.AddUser(user)

	if repo.Users[2].FirstName != "John" {
		t.Errorf("user not found on wanted id: %v", repo.Users[2])
	}
}

func TestDeleteUser(t *testing.T) {
	repo := CreateRepo()

	repo.DeleteUser(1)

	_, exists := repo.Users[1]

	if exists {
		t.Error("user still exists")
	}
}

func TestUpdateUser(t *testing.T) {
	repo := CreateRepo()

	updates := models.User{
		FirstName: "John",
	}

	err := repo.UpdateUser(updates, 1)

	if err != nil {
		t.Errorf("returned error: %v", err)
	}

	if repo.Users[1].FirstName != "John" {
		t.Errorf("User not updated: %v", repo.Users[1])
	}
}
