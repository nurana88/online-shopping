package main

import (
	"testing"

	config "github.com/nurana88/online-shopping/config"
)

func TestUserValidation(t *testing.T) {

	testUser := config.User{
		Name:     "Jeyn",
		Lastname: "McGregory",
		Password: "tErw123>",
	}
	if testUser.Validate() != nil {
		t.Error("Test failed:", testUser)
	}
}

