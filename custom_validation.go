package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/nurana88/online-shopping/pkg/domain"
)

func init() {
	addCustomValidators()
}

func addCustomValidators() {
	govalidator.TagMap["domain-country"] = func(s string) bool {
		_, err := domain.NewCountry(s)
		return err == nil
	}
}
