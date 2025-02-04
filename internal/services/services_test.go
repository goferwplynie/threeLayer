package services

import (
	"projekt/internal/models"
	"testing"
)

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
