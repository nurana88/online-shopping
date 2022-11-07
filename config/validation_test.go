package config

//
//import (
//	"errors"
//	"testing"
//)
//
//func TestCreateShouldFailIfUsernameIsSmallerThan3(t *testing.T) {
//	var user User
//
//	user.Id=123
//	user.Name=""
//
//	inserterMock := UserInserterMock{insertFunc: func(user User) error {
//		t.Fatal("this is not supposed to be called")
//		return nil
//	}}
//
//	err := NewUserCreate(inserterMock).CreateUser(user)
//
//	if err == nil {
//		t.Fatal("expected an erro for username lenght smaller than 3")
//	}
//}
//
//func TestCreateShouldCreateValidUsers(t *testing.T) {
//	user := User{123, "Stefano"}
//
//	mockCalled := false
//	inserterMock := UserInserterMock{insertFunc: func(user User) error {
//		mockCalled = true
//		return nil
//	}}
//
//	err := NewCreateUserUsecase(inserterMock).Create(user)
//
//	if err != nil {
//		t.Fatal("no error was expected")
//	}
//
//	if mockCalled == false {
//		t.Fatal("the user was supposed to be created")
//	}
//}
//
//func TestShouldRTeturnAnErrorWhenUserInserterFails(t *testing.T) {
//	user := User{"123", "Stefano"}
//
//	inserterMock := UserInserterMock{insertFunc: func(user User) error {
//		return errors.New("failed to insert")
//	}}
//
//	err := NewCreateUserUsecase(inserterMock).Create(user)
//	if err == nil {
//		t.Fatal("error was expected")
//	}
//
//	e := err.Error()
//
//	if e != "failed to insert" {
//		t.Fatal("was expecting the error from the inserter")
//	}
//}
//
