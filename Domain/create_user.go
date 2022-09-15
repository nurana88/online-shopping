package domain

import (
	"errors"
	"fmt"
	config "github.com/nurana88/online-shopping/config"
)

type createUser struct {
	dbActions DBActions
}

// Constructor
func NewUserCreate(dbActions DBActions) *createUser {
	return &createUser{dbActions: dbActions}
}

// ------* Create user in DB *------- //
func (c *createUser) CreateUser(user config.User) error {

	fmt.Println("user is in creat func...", user)
	if err := user.Validate(); err != nil {
		fmt.Println("Error in createUser", err)
		return errors.New("error in creating new user")
	}
	fmt.Println("User validated in Create func")
	if err := c.dbActions.InsertUser(user); err != nil {
		fmt.Println("error in Inserting user", err)
		return errors.New("error in inserting user")
	}

	fmt.Println("Created user after inserting...", user)

	return nil
}
