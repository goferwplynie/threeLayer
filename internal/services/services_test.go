package services

import (
	"projekt/internal/models"
	"projekt/internal/repository"
	"testing"
)

func CreateRepo() *repository.Repository {
	user := models.User{
		Id:        1,
		FirstName: "Dawid",
		LastName:  "Markiewicz",
		Age:       16,
		Group:     "admin",
	}

	return &repository.Repository{
		Users: map[int]models.User{
			1: user,
		},
		NextId: 2,
	}
}

func TestAgeCalculation(t *testing.T) {
	birthYear := 2000
	user := models.Request{
		BirthYear: &birthYear,
	}
	expectedAge := 25

	calculatedAge := CalculateAge(user)
	if calculatedAge != expectedAge {
		t.Errorf("expected %v but got %v", expectedAge, calculatedAge)
	}
}

func TestGroupCheck(t *testing.T) {
	group := "admin"
	user := models.Request{
		Group: &group,
	}
	if !CheckUserGroup(user) {
		t.Errorf("%v expected true got false", group)
	}
	group = "premium"
	user.Group = &group
	if !CheckUserGroup(user) {
		t.Errorf("%v expected true got false", group)
	}
	group = "user"
	user.Group = &group
	if !CheckUserGroup(user) {
		t.Errorf("%v expected true got false", group)
	}
	group = "adawds"
	user.Group = &group
	if CheckUserGroup(user) {
		t.Errorf("%v expected false got true", group)
	}
}

func CreateExampleUserRequest() *models.Request {
	firstName := "Dawid"
	lastName := "Markiewicz"
	birthYear := 2009
	group := "admin"

	return &models.Request{
		FirstName: &firstName,
		LastName:  &lastName,
		BirthYear: &birthYear,
		Group:     &group,
	}
}

func TestAddUser(t *testing.T) {
	repo := repository.New()
	bLayer := New(repo)
	user := CreateExampleUserRequest()

	err := bLayer.AddUser(*user)

	if err != nil {
		t.Errorf("error encountered: %v", err)
	}
}

func TestRaiseWhenNotAllFieldsProvided(t *testing.T) {
	repo := repository.New()
	bLayer := New(repo)
	user := CreateExampleUserRequest()
	user.FirstName = nil

	err := bLayer.AddUser(*user)

	if err == nil {
		t.Error("should throw error")
	}
}

func TestRaiseWhenWrongGroupProvided(t *testing.T) {
	repo := repository.New()
	bLayer := New(repo)
	user := CreateExampleUserRequest()
	group := "adawds"
	user.Group = &group

	err := bLayer.AddUser(*user)

	if err == nil {
		t.Error("should throw error")
	}
}

func TestUpdateUser(t *testing.T) {
	repo := CreateRepo()
	bLayer := New(repo)
	user := CreateExampleUserRequest()
	birthYear := 2010

	user.BirthYear = &birthYear
	user.FirstName = nil
	user.LastName = nil
	user.Group = nil

	err := bLayer.UpdateUser(*user, 1)
	if err != nil {
		t.Errorf("error thrown: %v", err)
	}

	if repo.Users[1].Age != 15 {
		t.Errorf("Expected age 15 but got %v instead", repo.Users[1].Age)
	}

}
