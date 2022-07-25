package domain

import (
	"errors"
	"github.com/nurana88/online-shopping/config"
	"testing"
)

// go test command runs the tests
// go test -bench=. runs benchmarks

type UserFinder struct {
}

func (u UserFinder) FindUserInDB(user config.User) error {
	return nil
}

func TestNewGetUser(t *testing.T) {

	testReqUser := config.LoginDetails{
		Email:    "jen@hf.com",
		Password: "tErw123>",
	}

	user, err := NewGetUserUsecase(UserFinder{}).GetUser(testReqUser)

	if err != nil {
		t.Fatalf("Failed to create user %v, an error occurred: %s", user, err.Error())
	}
}

// ---- NEW ---- //
type UserInserterMock struct {
	insertFunc func(user User) error
}

func (u UserInserterMock) Insert(user User) error {
	return u.insertFunc(user)
}

func TestCreateShouldFailIfUsernameIsSmallerThan3(t *testing.T) {
	user := User{"123", ""}

	inserterMock := UserInserterMock{insertFunc: func(user User) error {
		t.Fatal("this is not supposed to be called")
		return nil
	}}

	err := NewCreateUserUsecase(inserterMock).Create(user)

	if err == nil {
		t.Fatal("expected an erro for username lenght smaller than 3")
	}
}

func TestCreateShouldCreateValidUsers(t *testing.T) {
	user := User{"123", "Stefano"}

	mockCalled := false
	inserterMock := UserInserterMock{insertFunc: func(user User) error {
		mockCalled = true
		return nil
	}}

	err := NewCreateUserUsecase(inserterMock).Create(user)

	if err != nil {
		t.Fatal("no error was expected")
	}

	if mockCalled == false {
		t.Fatal("the user was supposed to be created")
	}
}

func TestShouldRTeturnAnErrorWhenUserInserterFails(t *testing.T) {
	user := User{"123", "Stefano"}

	inserterMock := UserInserterMock{insertFunc: func(user User) error {
		return errors.New("failed to insert")
	}}

	err := NewCreateUserUsecase(inserterMock).Create(user)
	if err == nil {
		t.Fatal("error was expected")
	}

	e := err.Error()

	if e != "failed to insert" {
		t.Fatal("was expecting the error from the inserter")
	}
}
```