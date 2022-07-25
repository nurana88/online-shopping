package domain

import (
	"testing"

	config "github.com/nurana88/online-shopping/config"
)

type UserInserter struct {
}

func (u UserInserter) InsertUser(user config.User) error {
	return nil
}

func TestNewCreateUser(t *testing.T) {

	testUser := config.User{
		Name:     "Jeyn",
		Lastname: "McGregory",
		Password: "tErw123>",
	}

	user, err := NewCreateUserUsecase(UserInserter{}).CreateUser(testUser)

	if err != nil {
		t.Fatalf("Failed to create user %v, an error occurred: %s", user, err.Error())
	}
}
