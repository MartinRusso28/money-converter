package repository

import (
	"github.com/MartinRusso28/money-converter/pkg/converter/model"
)

//Repository for storing data.
type Repository struct {
	Money    MoneyRepository
}

//Health return if the repository its available.
type Health interface {
	Check() error
}

//MoneyRepository define an interface for handling the money model persistence.
type MoneyRepository interface {
	GetRate(exchangeName string) (float64, error)
	GetEnabledExchanges() []string
	AddCurrency(currency model.Currency) error
}
