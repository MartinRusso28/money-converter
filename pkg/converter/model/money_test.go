package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fail_money_with_amount_less_than_0(t *testing.T) {
	money := Money{
		Amount: -1,
		Currency: Currency{
			CurrencyCode: "USD",
		},
	}

	valid := money.Valid()

	assert.False(t, valid)
}

func Test_can_create_money_with_valid_fields(t *testing.T) {
	money := Money{
		Amount: 1,
		Currency: Currency{
			CurrencyCode: "USD",
		},
	}

	valid := money.Valid()

	assert.True(t, valid)
}