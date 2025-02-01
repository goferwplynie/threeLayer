package models

type Request struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	BirthYear *int    `json:"birthYear"`
	Group     *string `json:"group"`
}
