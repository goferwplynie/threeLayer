package repository

import (
	"projekt/internal/models"
)

type Repository struct {
	Users map[int]models.User `json:"users"`
}

func (r *Repository) AddUser(user models.User) {
	r.Users[user.Id] = user
}

func (r *Repository) GetLastId() int {

	return 0
}
