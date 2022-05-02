package Usercases

import (
	"unicode"
)

type User struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string
	DateCreated    string `json:"date_created"`
}

type LoginDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Validate() *ApiErr {

	pwdSpace, pwdHasLower, pwdHasUpper, pwdHasNum, pwdHasChar, pwdLength := false, false, false, false, false, false

	// Check if name is only letters
	for _, char := range user.Name {
		if unicode.IsLetter(char) == false {
			return NewBadRequest("Incorrect Name")
		}
	}

	// Check if lastname is only letters
	for _, char := range user.Lastname {
		if unicode.IsLetter(char) == false {
			return NewBadRequest("Incorrect Lastname")
		}
	}

	// Check password for lower, upper, number, character, non-space and length cases
	for _, char := range user.Password {
		switch {
		case unicode.IsLower(char):
			pwdHasLower = true
		case unicode.IsUpper(char):
			pwdHasUpper = true
		case unicode.IsNumber(char):
			pwdHasNum = true
		case unicode.IsSymbol(char):
			pwdHasChar = true
		case unicode.IsSpace(char):
			pwdSpace = true
		}
	}

	if len(user.Password) >= 6 {
		pwdLength = true
	}

	if pwdSpace || !pwdHasLower || !pwdHasUpper || !pwdHasNum || !pwdHasChar || !pwdLength {
		return NewBadRequest("Incorrect Password")
	}

	// Check if passwords are matching
	if user.Password != user.PasswordRepeat {
		return NewBadRequest("Passwords are not matching")
	}

	// Check Email

	return nil
}
