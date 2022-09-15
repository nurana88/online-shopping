package domain

import "github.com/nurana88/online-shopping/config"

type DBActions interface {
	InsertUser(user config.User) error
	FindUserInDB(user config.User) error
}

type UserCreateAction interface {
	CreateUser(user config.User) error
}

type UserGetAction interface {
	GetUser(request config.LoginDetails) (*config.User, error)
}
