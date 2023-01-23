// Package domain adds new type of Country as a string
package domain

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Country to represent country
type Country string

var (
	countryPattern = `^[a-zA-Z]{2}$`
	isIso2Code     = regexp.MustCompile(countryPattern).MatchString

	// CountryNZ is NZ country object
	CountryNZ, _ = NewCountry("nz")
)

// NewCountry creates new hellofresh country from iso2code.
//
// Returns an ErrInvalidCountryFormat error if input string does not match iso2Code rules.
func NewCountry(iso2Code string) (Country, error) {
	if !isIso2Code(iso2Code) {
		return "", fmt.Errorf("Invalid country format. Should be of type iso2code, e.g.:\"us\", received '%s'", iso2Code)
	}

	return Country(strings.ToLower(iso2Code)), nil
}

// UnmarshalJSON enforces lowercase value
func (c *Country) UnmarshalJSON(data []byte) error {
	var iso2Code string
	if err := json.Unmarshal(data, &iso2Code); err != nil {
		return err
	}
	v, err := NewCountry(iso2Code)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

// Is compare two countries together
func (c *Country) Is(country *Country) bool {
	return *c == *country
}

// CountryConverter is a custom country converter for gorilla `schema`
func CountryConverter(value string) reflect.Value {
	if c, err := NewCountry(value); err == nil {
		return reflect.ValueOf(c)
	}

	return reflect.Value{}
}
