package domain

import (
	"fmt"

	config "github.com/nurana88/online-shopping/config"
)

func (c *DBUserUsercase) GetUser(request config.LoginDetails) (*config.User, error) {

	reqPassword := config.GetMd5(request.Password)

	userRequest := config.User{
		Email:    request.Email,
		Password: reqPassword,
	}

	if err := c.dbActions.FindUserInDB(userRequest); err != nil {
		return nil, err
	}

	fmt.Println("User was found...", userRequest)

	return &userRequest, nil
}
