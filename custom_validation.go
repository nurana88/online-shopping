package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/nurana88/online-shopping/pkg/local"
)

func init() {
	addCustomValidators()
}

func addCustomValidators() {
	govalidator.TagMap["local-country"] = func(s string) bool {
		_, err := local.NewCountry(s)
		return err == nil
	}
}
