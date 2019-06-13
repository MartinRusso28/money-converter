package model

import (
	"github.com/asaskevich/govalidator"
)

//Currency contains CurrencyCode and Rate of a Currency.
type Currency struct {
	CurrencyCode string  `valid:"required, length(3|3)"`
	Rate         float64 `json:"-" `
}

//Valid currency struct.
func (c Currency) Valid() bool {
	valid, _ := govalidator.ValidateStruct(c)
	return valid
}
