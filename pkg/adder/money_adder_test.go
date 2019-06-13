package moneyadder

import(
	"testing"
	"github.com/MartinRusso28/money-converter/pkg/converter"
	"github.com/MartinRusso28/money-converter/pkg/converter/model"
	"github.com/stretchr/testify/assert"
	"money-converter/test"
)

func Test_api_can_add_with_valid_money(t *testing.T) {

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

func Test_repo_can_add_with_valid_money(t *testing.T) {

	repo := mock.GetRepository()

	pesoC := model.Currency{
		CurrencyCode: "ARS",
		Rate: 10,
	}

	dolarC := model.Currency{
		CurrencyCode: "USD",
		Rate: 40,
	}

	err := repo.Money.AddCurrency(pesoC)
	err = repo.Money.AddCurrency(dolarC)

	moneyAdder := MoneyAdder{
		CurrencyConverter: moneyconverter.CurrencyConverter{
			RateStrategy: moneyconverter.RepositoryRateStrategy{Repo: repo},
			ToCurrency: dolarC,
		},
	}

	dolar := model.Currency{
		CurrencyCode:"USD",
	}

	peso := model.Currency{
		CurrencyCode:"ARS",
	}

	money1 := model.Money{
		Amount: 10,
		Currency: dolar,
	}

	money2 := model.Money{
		Amount: 100,
		Currency: peso,
	}

	moneyAdder.AddMoney(money1)
	moneyAdder.AddMoney(money2)

	result, err := moneyAdder.CalculateSumResult()

	assert.Greater(t, result, float64(0))
	assert.NoError(t, err)
}

