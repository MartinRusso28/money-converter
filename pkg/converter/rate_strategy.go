package moneyconverter

import (
	"money-converter/pkg/rates"
	r "money-converter/internal/repository"
)


//RaterStrategy is an interface that define different ways of get the rates of the exchanges.
type RaterStrategy interface {
	GetRate(fromCode string) (float64, error)
}

//APIRateStrategy calcualte the rate using an external API.
type APIRateStrategy struct {
}

//GetRate return the rate of a exchange using an external API.
func (api APIRateStrategy) GetRate(fromCode string) (float64, error) {
	newRate, err := ratescalc.GetExchangeRate(fromCode)

	if err != nil {
		return -1, err
	}

	return newRate, nil
}

//RepositoryRateStrategy calculate the rate using a repository.
type RepositoryRateStrategy struct {
	Repo r.Repository
}

//GetRate return the rate of a exchange using a repository.
func (repoSt RepositoryRateStrategy) GetRate(fromCode string) (float64, error) {
	newRate, err := repoSt.Repo.Money.GetRate(fromCode)

	if err != nil {
		return -1, err
	}

	return newRate, nil
}
