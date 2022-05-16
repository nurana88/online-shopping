package domain

import (
	"errors"
	"fmt"

	infra "github.com/nurana88/online-shopping/Infrastructure"
	config "github.com/nurana88/online-shopping/config"
)

// ------* Register user *------- //

type DBUserUsercase struct {
	dbActions *infra.DbActions
}

func NewDbUserUsercase(dbActions *infra.DbActions) *DBUserUsercase {
	return &DBUserUsercase{dbActions: dbActions}
}

func (c *DBUserUsercase) CreateUser(user config.User) (*config.User, error) {

	if err := user.Validate(); err != nil {
		fmt.Println("Error in createUser", err)
		return nil, errors.New("Error in creating new user")
	}
	c.dbActions.InsertUser(user)
	fmt.Println("Created user after inserting...", user)

	return &user, nil
}
