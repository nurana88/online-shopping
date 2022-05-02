package Usercases

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(request LoginDetails) (*User, *ApiErr) {

	reqPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, NewInternalError("Cant hash password of request")
	}

	userRequest := &User{
		Email:    request.Email,
		Password: string(reqPassword),
	}

	if err := userRequest.FindUserInDB(); err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRequest.Password), []byte(request.Password))
	if err != nil {
		fmt.Println("Err in comparehash is: ", err)
		return nil, NewBadRequest("Password is not correcT")
	}
	fmt.Println("User was found...", userRequest)

	return userRequest, nil
}
