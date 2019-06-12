package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fail_currency_without_CurrencyCode(t *testing.T) {
	currency := Currency{
	}

	valid := currency.Valid()

	assert.False(t, valid)
}

func Test_fail_currency_with_CurrencyCode_length_greater_than_3(t *testing.T) {
	currency := Currency{
		CurrencyCode: "AAAA",
	}

	valid := currency.Valid()

	assert.False(t, valid)
}

func Test_fail_currency_with_CurrencyCode_length_less_than_3(t *testing.T) {
	currency := Currency{
		CurrencyCode: "AA",
	}

	valid := currency.Valid()

	assert.False(t, valid)
}

func Test_can_create_currency_with_valid_fields(t *testing.T) {
	currency := Currency{
		CurrencyCode: "USD",
	}

	valid := currency.Valid()

	assert.True(t, valid)
}