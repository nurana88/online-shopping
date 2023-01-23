package domain

import (
	"testing"

	config "github.com/nurana88/online-shopping/config"
)

func TestNewCreateUser(t *testing.T) {
	var dbActions DBActions

	testUser := config.User{
		Name:           "Jeyn",
		Lastname:       "McGregory",
		Password:       "tErw123<>@",
		PasswordRepeat: "tErw123<>@",
	}

	err := NewUserCreate(dbActions).CreateUser(testUser)

	if err != nil {
		t.Fatalf("Failed to create user %v, an error occurred: %s", testUser, err.Error())
	}
}
