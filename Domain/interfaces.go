package domain

import "github.com/nurana88/online-shopping/config"

type UserDBInsert interface {
	InsertUser(user config.User) error
}

type UserDBFind interface {
	FindUserInDB(user config.User) error
}
