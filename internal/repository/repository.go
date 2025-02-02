package repository

import (
	"errors"
	"projekt/internal/models"
)

type Repository struct {
	Users  map[int]models.User `json:"users"`
	NextId int
}

func New() *Repository {
	return &Repository{
		Users:  make(map[int]models.User),
		NextId: 1,
	}
}

func (r *Repository) AddUser(user models.User) {
	user.Id = r.NextId
	r.Users[user.Id] = user
	r.NextId++
}

func (r *Repository) GetAll() []models.User {
	users := make([]models.User, len(r.Users))

	for _, v := range r.Users {
		users = append(users, v)
	}

	return users
}

func (r *Repository) GetUserById(id int) (models.User, error) {
	user, exists := r.Users[id]

	if !exists {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *Repository) DeleteUser(id int) error {
	_, exists := r.Users[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(r.Users, id)

	return nil
}
