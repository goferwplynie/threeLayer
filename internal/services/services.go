package services

import (
	"errors"
	"projekt/internal/models"
	"projekt/internal/repository"
	"reflect"
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
	return slices.Contains(availableGroups, *user.Group)
}

func CalculateAge(user models.Request) int {
	currentYear := time.Now().Year()

	age := currentYear - *user.BirthYear

	return age
}

func CheckAllFieldsProvided(userRequest models.Request) bool {
	userRequestFields := reflect.ValueOf(userRequest)
	for i := 0; i < userRequestFields.NumField(); i++ {
		field := userRequestFields.Field(i)
		if field.IsNil() {
			return false
		}
	}
	return true
}

func (b *BuisnessLayer) AddUser(userRequest models.Request) error {
	if !CheckAllFieldsProvided(userRequest) {
		return errors.New("provide all fields")
	}

	if !CheckUserGroup(userRequest) {
		return errors.New("wrong group provided")
	}

	user := models.User{
		FirstName: *userRequest.FirstName,
		LastName:  *userRequest.LastName,
		Age:       CalculateAge(userRequest),
		Group:     *userRequest.Group,
	}
	b.repo.AddUser(user)

	return nil
}

func (b *BuisnessLayer) GetAllUsers() []models.User {
	return b.repo.GetAll()
}

func (b *BuisnessLayer) GetUser(id int) (models.User, error) {
	return b.repo.GetUserById(id)
}

func (b *BuisnessLayer) DeleteUser(id int) error {
	return b.repo.DeleteUser(id)
}

func (b *BuisnessLayer) UpdateUser(userRequest models.Request, id int) error {
	user := models.User{}

	if userRequest.Group != nil {
		if !CheckUserGroup(userRequest) {
			return errors.New("wrong group provided")
		}
		user.Group = *userRequest.Group
	}
	if userRequest.BirthYear != nil {
		user.Age = CalculateAge(userRequest)
	}
	if userRequest.FirstName != nil {
		user.FirstName = *userRequest.FirstName
	}
	if userRequest.LastName != nil {
		user.LastName = *userRequest.LastName
	}

	err := b.repo.UpdateUser(user, id)
	return err
}
