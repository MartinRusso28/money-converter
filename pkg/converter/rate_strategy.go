package moneyconverter

import (
	"money-converter/pkg/rates"
	r "money-converter/internal/repository"
)


//RateStrategy -
type RateStrategy interface {
	GetRate(fromCode string) (float64, error)
}

//APIRateStrategy -
type APIRateStrategy struct {
}

//GetRate -
func (api APIRateStrategy) GetRate(fromCode string) (float64, error) {
	newRate, err := ratescalc.GetExchangeRate(fromCode)

	if err != nil {
		return -1, err
	}

	return newRate, nil
}

//RepositoryRateStrategy -
type RepositoryRateStrategy struct {
	repo r.Repository
}

//GetRate -
func (repoSt RepositoryRateStrategy) GetRate(fromCode string) (float64, error) {
	newRate, err := repoSt.repo.Money.GetRate(fromCode)

	if err != nil {
		return -1, err
	}

	return newRate, nil
}
