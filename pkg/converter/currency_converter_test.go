package moneyconverter

import (
	"github.com/MartinRusso28/money-converter/pkg/converter/model"
	"money-converter/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_api_strategy_cant_convert_with_unexisting_to_currency_code(t *testing.T) {

	dolar := model.Currency{
		CurrencyCode: "USD",
	}

	dolarMoney := model.Money{
		Amount:   10,
		Currency: dolar,
	}

	mexican := model.Currency{
		CurrencyCode: "QQQ",
	}

	currencyConverter := CurrencyConverter{
		RateStrategy: APIRateStrategy{},
		ToCurrency:         mexican,
	}

	_, err := currencyConverter.Convert(&dolarMoney)

	assert.Error(t, err)
}

func Test_api_strategy_cant_convert_with_unexisting_from_currency_code(t *testing.T) {

	dolar := model.Currency{
		CurrencyCode: "QQQ",
	}

	dolarMoney := model.Money{
		Amount:   10,
		Currency: dolar,
	}

	mexican := model.Currency{
		CurrencyCode: "USD",
	}

	currencyConverter := CurrencyConverter{
		RateStrategy: APIRateStrategy{},
		ToCurrency:         mexican,
	}

	_, err := currencyConverter.Convert(&dolarMoney)

	assert.Error(t, err)
}

func Test_api_strategy_can_convert_with_existing_currency_codes(t *testing.T) {

	dolar := model.Currency{
		CurrencyCode: "USD",
	}

	dolarMoney := model.Money{
		Amount:   10,
		Currency: dolar,
	}

	mexican := model.Currency{
		CurrencyCode: "MXN",
	}

	currencyConverter := CurrencyConverter{
		RateStrategy: APIRateStrategy{},
		ToCurrency:         mexican,
	}

	convertedMoney, err := currencyConverter.Convert(&dolarMoney)

	assert.NoError(t, err)
	assert.Greater(t, convertedMoney.Amount, float64(0))
	assert.NotNil(t, convertedMoney)
	assert.NotNil(t, convertedMoney.CurrencyCode)
}


func Test_repo_strategy_cant_convert_with_unexisting_from_currency_code(t *testing.T) {

	repoMock := mock.GetRepository()

	currency := model.Currency{
		CurrencyCode: "USD",
		Rate: 2,
	}

	err := repoMock.Money.AddCurrency(currency)

	dolar := model.Currency{
		CurrencyCode: "USD",
	}

	dolarMoney := model.Money{
		Amount:   10,
		Currency: dolar,
	}

	mexican := model.Currency{
		CurrencyCode: "MXN",
	}

	currencyConverter := CurrencyConverter{
		RateStrategy: RepositoryRateStrategy{repo: repoMock},
		ToCurrency:         mexican,
	}

	_, err = currencyConverter.Convert(&dolarMoney)

	assert.Error(t, err)
}


func Test_repo_strategy_cant_convert_with_unexisting_to_currency_code(t *testing.T) {

	repoMock := mock.GetRepository()

	currency := model.Currency{
		CurrencyCode: "USD",
		Rate: 2,
	}

	err := repoMock.Money.AddCurrency(currency)

	dolar := model.Currency{
		CurrencyCode: "MXN",
	}

	dolarMoney := model.Money{
		Amount:   10,
		Currency: dolar,
	}

	mexican := model.Currency{
		CurrencyCode: "USD",
	}

	currencyConverter := CurrencyConverter{
		RateStrategy: RepositoryRateStrategy{repo: repoMock},
		ToCurrency:         mexican,
	}

	_, err = currencyConverter.Convert(&dolarMoney)

	assert.Error(t, err)
}


func Test_repo_strategy_can_convert_with_existing_currency_codes(t *testing.T) {

	repoMock := mock.GetRepository()

	dolar := model.Currency{
		CurrencyCode: "USD",
		Rate:         1,
	}

	dolarMoney := model.Money{
		Amount:   10,
		Currency: dolar,
	}

	mexican := model.Currency{
		CurrencyCode: "MXN",
		Rate:         2,
	}

	err := repoMock.Money.AddCurrency(dolar)
	err = repoMock.Money.AddCurrency(mexican)

	currencyConverter := CurrencyConverter{
		RateStrategy: RepositoryRateStrategy{repo: repoMock},
		ToCurrency:         mexican,
	}

	convertedMoney, err := currencyConverter.Convert(&dolarMoney)

	if err != nil {
		panic(err)
	}

	assert.NoError(t, err)
	assert.Greater(t, convertedMoney.Amount, float64(0))
	assert.NotNil(t, convertedMoney)
	assert.NotNil(t, convertedMoney.CurrencyCode)
}


