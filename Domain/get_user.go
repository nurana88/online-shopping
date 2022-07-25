package domain

import (
	"fmt"
	config "github.com/nurana88/online-shopping/config"
)

type UserFind interface {
	FindUserInDB(user config.User) error
}

type GetUserUsecase struct {
	userRepo UserFind
}

func NewGetUserUsecase(userRepo UserFind) *GetUserUsecase {
	return &GetUserUsecase{userRepo: userRepo}
}

func (g *GetUserUsecase) GetUser(request config.LoginDetails) (*config.User, error) {

	reqPassword := config.GetMd5(request.Password)

	userRequest := config.User{
		Email:    request.Email,
		Password: reqPassword,
	}

	if err := g.userRepo.FindUserInDB(userRequest); err != nil {
		return nil, err
	}

	fmt.Println("User was found...", userRequest)

	return &userRequest, nil
}
