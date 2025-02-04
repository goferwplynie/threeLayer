package services

import (
	"errors"
	"projekt/internal/models"
	"projekt/internal/repository"
	"slices"
	"time"
)

type BuisnessLayer struct {
	repo *repository.Repository
}

var availableGroups = []string{"admin", "premium", "user"}

func New(repo *repository.Repository) *BuisnessLayer {
	return &BuisnessLayer{
		repo: repo,
	}
}

func CheckUserGroup(user models.Request) bool {
	if user.Group == nil || !slices.Contains(availableGroups, *user.Group) {
		return false
	}

	return true
}

func CalculateAge(user models.Request) int {
	currentYear := time.Now().Year()

	age := currentYear - *user.BirthYear

	return age
}

func MapRequestToUser(userRequest models.Request) (models.User, error) {
	user := models.User{}
	if userRequest.FirstName == nil {
		return models.User{}, errors.New("no firstName provided")
	}
	user.FirstName = *userRequest.FirstName
	if userRequest.LastName == nil {
		return models.User{}, errors.New("no lastName provided")
	}
	user.LastName = *userRequest.LastName
	if userRequest.BirthYear == nil {
		return models.User{}, errors.New("no birthYear provided")
	}
	user.Age = CalculateAge(userRequest)
	if !CheckUserGroup(userRequest) {
		return models.User{}, errors.New("no group or wrong group provided")
	}
	user.Group = *userRequest.Group

	return user, nil
}
