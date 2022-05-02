package Usercases

import (
	"fmt"
)

// ------* Register user *------- //
/*
- Name and lastname should contain only letters
- Passwords should contain uppercase, lowercase letter, number and special character
- Password should be greater than 6 characters
*/
func CreateUser(user User) (*User, *ApiErr) {

	if err := user.Validate(); err != nil {
		fmt.Println("Error in createUser", err)
		return nil, NewBadRequest("Error in creating new user")
	}
	user.InsertUser()
	fmt.Println("Created user after inserting...", user)

	return &user, nil
}
