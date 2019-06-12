package network

import (
	"github.com/asaskevich/govalidator"
)

//Response represent a request response.
type Response struct {
	StatusCode int `valid:"required"`
	Body       interface{}
	Error      error
}

//Valid response struct.
func (r Response) Valid() bool {
	valid, _ := govalidator.ValidateStruct(r)
	return valid
}