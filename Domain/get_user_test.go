package domain

import (
	"github.com/nurana88/online-shopping/config"
	"testing"
)

// go test command runs the tests
// go test -bench=. runs benchmarks

func TestNewGetUser(t *testing.T) {
	var dbActions DBActions

	testReqUser := config.LoginDetails{
		Email:    "jen@domain.com",
		Password: "tErw123>",
	}

	user, err := NewGetUser(dbActions).GetUser(testReqUser)

	if err != nil {
		t.Fatalf("Failed to create user %v, an error occurred: %s", user, err.Error())
	}
}

// ---- NEW ---- //
type UserInserterMock struct {
	insertFunc func(user config.User) error
}

func (u UserInserterMock) Insert(user config.User) error {
	return u.insertFunc(user)
}
