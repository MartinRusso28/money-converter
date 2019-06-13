package model

import (
	"github.com/asaskevich/govalidator"
)

//Money -
type Money struct {
	Amount float64
	Currency
}

//Valid money struct.
func (m Money) Valid() bool {
	valid, _ := govalidator.ValidateStruct(m)
	return valid && m.Amount > 0
}
