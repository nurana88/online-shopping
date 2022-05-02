package Validation

import (
	"unicode"

	"github.com/nurana88/shop-template/Domain"
	"github.com/nurana88/shop-template/Domain/Usercases"
)

func (user *Domain.User) Validate() *Usercases.ApiErr {

	pwdNoSpace, pwdHasLower, pwdHasUpper, pwdHasNum, pwdHasChar, pwdLength := true, true, true, true, true, true

	// Check if name is only letters
	for _, char := range user.Name {
		if unicode.IsLetter(char) == false {
			return Usercases.NewBadRequest("Error in creating user name")
		}
	}

	// Check if lastname is only letters
	for _, char := range user.Lastname {
		if unicode.IsLetter(char) == false {
			return Usercases.NewBadRequest("Error in creating user lastname")
		}
	}

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
			pwdNoSpace = false
		}
	}

	if len(user.Password) >= 6 {
		pwdLength = true
	}

	if !pwdNoSpace || !pwdHasLower || !pwdHasUpper || !pwdHasNum || !pwdHasChar || !pwdLength {
		return Usercases.NewBadRequest("Error in creating user password")
	}

	return nil
}
