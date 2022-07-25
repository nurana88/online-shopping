package domain

import (
	"errors"
	"fmt"

	config "github.com/nurana88/online-shopping/config"
)

// ------* Register user *------- //
type UserInsert interface {
	InsertUser(user config.User) error
}
type CreateUserUsecase struct {
	userRepo UserInsert
}

// Constructor
func NewCreateUserUsecase(userRepo UserInsert) *CreateUserUsecase {
	return &CreateUserUsecase{userRepo: userRepo}
}

func (c *CreateUserUsecase) CreateUser(user config.User) (*config.User, error) {

	if err := user.Validate(); err != nil {
		fmt.Println("Error in createUser", err)
		return nil, errors.New("error in creating new user")
	}
	if err := c.userRepo.InsertUser(user); err != nil {
		fmt.Println("error in Inserting user", err)
		return nil, errors.New("error in inserting user")
	}

	fmt.Println("Created user after inserting...", user)

	return &user, nil
}
