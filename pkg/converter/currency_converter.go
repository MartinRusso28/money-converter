package moneyconverter

import (
	"money-converter/pkg/converter/model"
)

//CurrencyConverter -
type CurrencyConverter struct {
	RateStrategy RateStrategy
	ToCurrency          model.Currency
}

//Convert -
func (cc CurrencyConverter) Convert(from *model.Money) (*model.Money, error) {
	fromRate, err := cc.RateStrategy.GetRate(from.CurrencyCode)

	if err != nil {
		return nil, err
	}

	toRate, err := cc.RateStrategy.GetRate(cc.ToCurrency.CurrencyCode)

	if err != nil {
		return nil, err
	}

	newRate :=  1 / fromRate * toRate
	newAmount := newRate * from.Amount

	resultMoney := model.Money{
		Amount:   newAmount,
		Currency: cc.ToCurrency,
	}

	return &resultMoney, nil
}