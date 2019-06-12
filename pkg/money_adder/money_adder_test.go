package moneyadder

import(
	"testing"
	"money-converter/pkg/converter"
	"money-converter/pkg/converter/model"
	"github.com/stretchr/testify/assert"
)

func Test_can_add_with_valid_money(t *testing.T) {

	japanese := model.Currency{
		CurrencyCode:"JPY",
	}

	moneyAdder := MoneyAdder{
		CurrencyConverter: moneyconverter.CurrencyConverter{
			RateStrategy: moneyconverter.APIRateStrategy{},
			ToCurrency: japanese,
		},
	}

	dolar := model.Currency{
		CurrencyCode:"USD",
	}

	mexican := model.Currency{
		CurrencyCode:"MXN",
	}

	money1 := model.Money{
		Amount: 10,
		Currency: dolar,
	}

	money2 := model.Money{
		Amount: 100,
		Currency: mexican,
	}

	money3 := model.Money{
		Amount: 1234,
		Currency: japanese,
	}

	moneyAdder.AddMoney(money1)
	moneyAdder.AddMoney(money2)
	moneyAdder.AddMoney(money3)

	result, err := moneyAdder.CalculateSumResult()

	assert.Greater(t, result, float64(0))
	assert.NoError(t, err)
}

