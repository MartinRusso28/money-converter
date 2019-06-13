package moneyadder

import(
	"money-converter/pkg/converter"
	"money-converter/pkg/converter/model"
)

//MoneyAdder is a struct used for sum different types of exchanges.
type MoneyAdder struct {
	moneyconverter.CurrencyConverter
	Monies []model.Money
}

//AddMoney to the moneyAdder.
func (adder *MoneyAdder) AddMoney(money model.Money){
	adder.Monies = append(adder.Monies, money)
}

//CalculateSumResult return the sum of all the monies that are in the money adder.
func (adder *MoneyAdder) CalculateSumResult() (float64, error) {
	var result float64

	for _, money := range adder.Monies {
		conversion, err := adder.Convert(&money)

		if err != nil {
			return -1, err
		}

		result += conversion.Amount
	}

	return result, nil
}

