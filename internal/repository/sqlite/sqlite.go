package sqlite

import (
	"money-converter/pkg/converter/model"
	r "money-converter/internal/repository"
	
	"errors"
	"github.com/jinzhu/gorm"
)

var (
	database *gorm.DB
)

//GetDatabase create a sqlite database and the all the business model.
func GetDatabase() r.Repository {
	var err error

	database, err = gorm.Open("sqlite3", "exchange.db")
	if err != nil {
		panic("Unable to open DB")
	}

	generateModels(database)

	repo := r.Repository{
		Money: MoneyRepository{},
	}

	return repo
}

func generateModels(database *gorm.DB) {

	database.LogMode(false)

	database.AutoMigrate(&model.Currency{})
	database.Create(&model.Currency{CurrencyCode: "USD", Rate: 1.5})
	database.Create(&model.Currency{CurrencyCode: "ARS", Rate: 40})
	database.Create(&model.Currency{CurrencyCode: "MXC", Rate: 20})

	database.Model(&model.Currency{}).AddUniqueIndex("idx_currency_name", "currency_code")
}

type sqliteHealth struct {
}

func (hc sqliteHealth) Check() error {
	return database.DB().Ping()
}

//MoneyRepository handle all the business logic of the exchanges.
type MoneyRepository struct {
}

//GetRate return a rate that is stored in the sqlite DB.
func (moneyRepo MoneyRepository) GetRate(exchangeName string) (float64, error) {
	var rates []model.Currency

	database.Where(&model.Currency{CurrencyCode: exchangeName}).Find(&rates)

	if len(rates) == 0{
		return -1, errors.New("Unexisting currency")
	}

	return rates[0].Rate, nil
}

//GetEnabledExchanges return all the rates that are stored in the sqlite DB.-
func (moneyRepo MoneyRepository) GetEnabledExchanges() ([]string) {
	var currencies []model.Currency
	var exchanges []string

	database.Select("currency_code").Find(&currencies)

	for _, v := range currencies {
		exchanges = append(exchanges, v.CurrencyCode)
	}

	return exchanges
}

//AddCurrency to the sqlite DB.
func (moneyRepo MoneyRepository) AddCurrency(currency model.Currency) error {
	err := database.Save(currency).GetErrors()

	return err[0]
}