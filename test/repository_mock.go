package mock

import (
	"github.com/MartinRusso28/money-converter/pkg/converter/model"
	"github.com/MartinRusso28/money-converter/internal/repository"
	"errors"
)

//MoneyRepositoryMock -
type MoneyRepositoryMock struct {
	currencies   []model.Currency
}

//GetRate -
func (repo MoneyRepositoryMock) GetRate(exchangeName string) (float64, error){
	for i := range repo.currencies {
		if repo.currencies[i].CurrencyCode == exchangeName {
			return repo.currencies[i].Rate, nil
		}
	}

	return -1, errors.New("unexisting exchange")
}

//GetEnabledExchanges -
func (repo *MoneyRepositoryMock) GetEnabledExchanges() []string {
	var exchanges []string

	for _, v := range repo.currencies {
		exchanges = append(exchanges, v.CurrencyCode)
	}

	return exchanges
}

//AddCurrency -
func (repo *MoneyRepositoryMock) AddCurrency(currency model.Currency) error {
	repo.currencies = append(repo.currencies, currency)
	return nil
}

//GetRepository -
func GetRepository() repository.Repository {
	return repository.Repository{
		Money:    &MoneyRepositoryMock{},
	}
}