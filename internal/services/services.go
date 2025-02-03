package services

import "projekt/internal/repository"

type BuisnessLayer struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *BuisnessLayer {
	return &BuisnessLayer{
		repo: repo,
	}
}
