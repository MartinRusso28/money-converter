package moneyadder

import(
	"money-converter/pkg/converter"
	"money-converter/pkg/converter/model"
)

//MoneyAdder -
type MoneyAdder struct {
	moneyconverter.CurrencyConverter
	Monies []model.Money
}

//AddMoney -
func (adder *MoneyAdder) AddMoney(money model.Money){
	adder.Monies = append(adder.Monies, money)
}

//CalculateSumResult -
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

