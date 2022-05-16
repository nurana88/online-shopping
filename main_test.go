package main

import (
	"testing"

	domain "github.com/nurana88/online-shopping/Domain"
)

func TestUserValidation(t *testing.T) {

	testUser := domain.User{
		Name:     "Jeyn",
		Lastname: "McGregory",
		Password: "tErw123>",
	}
	if testUser.Validate() != nil {
		t.Error("Test failed:", testUser)
	}
}
