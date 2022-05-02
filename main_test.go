package main

import (
	"testing"

	"github.com/nurana88/online-shopping/Domain/Usercases"
)

func TestUserValidation(t *testing.T) {

	testUser := Usercases.User{
		Name:     "Jeyn",
		Lastname: "McGregory",
		Password: "tErw123>",
	}
	if testUser.Validate() != nil {
		t.Error("Test failed:", testUser)
	}
}
