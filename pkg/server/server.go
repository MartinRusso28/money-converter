package moneysrv

import (
	"errors"
	"github.com/MartinRusso28/money-converter/internal/network"
	r "github.com/MartinRusso28/money-converter/internal/repository"
	"github.com/MartinRusso28/money-converter/pkg/converter/model"
	"github.com/MartinRusso28/money-converter/pkg/adder"
	"github.com/MartinRusso28/money-converter/pkg/converter"
	ratescalc "github.com/MartinRusso28/money-converter/pkg/rates"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetMainEngine return money converter's server.
func GetMainEngine(repo r.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	env := environment{repo: repo}

	router.GET("/exchanges/api/convert", env.setConvertParams, env.apiConvertExchange)
	router.GET("/exchanges/api", env.getAPIEnabledExchanges)
	router.POST("/exchanges/api/sum", env.getMonies, env.apiSumMoney)

	router.GET("/exchanges/repo/convert", env.setConvertParams, env.repoConvertExchange)
	router.GET("/exchanges/repo", env.getRepoEnabledExchanges)
	router.POST("/exchanges/repo/sum", env.getMonies, env.repoSumMoney)

	return router
}

type environment struct {
	repo r.Repository
}

func (env environment) getAPIEnabledExchanges(context *gin.Context) {
	enabledEx, err := ratescalc.GetEnabledExchanges()

	if err != nil {
		env.internalServerError(context, err)
	}

	env.respondEnabledExchanges(context, enabledEx)
}

func (env environment) getRepoEnabledExchanges(context *gin.Context) {
	enabledEx := env.repo.Money.GetEnabledExchanges()
	env.respondEnabledExchanges(context, enabledEx)
}

func (env environment) respondEnabledExchanges(context *gin.Context, exchanges []string) {
	response := network.Response{
		StatusCode: 200,
		Body:       exchanges,
	}

	env.respond(context, response)
}

func (env environment) setConvertParams(context *gin.Context) {
	fromCurrency := context.Query("from")
	toCurrency := context.Query("to")
	moneyAmount := context.Query("amount")

	moneyAmountFormatted, err := strconv.Atoi(moneyAmount)

	if err != nil || fromCurrency == "" || toCurrency == "" || moneyAmountFormatted <= 0 {
		env.badRequest(context)
		return
	}

	context.Set("from", fromCurrency)
	context.Set("to", toCurrency)
	context.Set("amount", moneyAmountFormatted)
}

func (env environment) apiConvertExchange(context *gin.Context) {
	toCurrency := model.Currency{
		CurrencyCode: context.GetString("to"),
	}

	amount := context.GetInt("amount")

	fromMoney := model.Money{
		Amount: float64(amount),
		Currency: model.Currency{
			CurrencyCode: context.GetString("from"),
		},
	}

	converter := moneyconverter.CurrencyConverter{
		RateStrategy: moneyconverter.APIRateStrategy{},
		ToCurrency:   toCurrency,
	}

	convertedMoney, err := converter.Convert(&fromMoney)

	if err != nil {
		env.internalServerError(context, err)
	}

	response := network.Response{
		StatusCode: 200,
		Body:       convertedMoney,
	}

	env.respond(context, response)
}

func (env environment) repoConvertExchange(context *gin.Context) {
	toCurrency := model.Currency{
		CurrencyCode: context.GetString("to"),
	}

	amount := context.GetInt("amount")

	fromMoney := model.Money{
		Amount: float64(amount),
		Currency: model.Currency{
			CurrencyCode: context.GetString("from"),
		},
	}

	converter := moneyconverter.CurrencyConverter{
		RateStrategy: moneyconverter.RepositoryRateStrategy{Repo: env.repo},
		ToCurrency:   toCurrency,
	}

	convertedMoney, err := converter.Convert(&fromMoney)

	if err != nil {
		env.internalServerError(context, err)
	}

	response := network.Response{
		StatusCode: 200,
		Body:       convertedMoney,
	}

	env.respond(context, response)
}

func (env environment) getMonies(context *gin.Context) {
	moniesRequest := SumParams{}
	err := context.BindJSON(&moniesRequest)

	if err != nil {
		panic(err)
	}

	for _, v := range moniesRequest.Monies {
		if !v.Valid() {
			env.badRequest(context)
			return
		}
	}

	context.Set("monies", moniesRequest.Monies)
	context.Set("toCurrency", moniesRequest.ToCurrency)
}

func (env environment) apiSumMoney(context *gin.Context) {
	monies, exists := context.Get("monies")
	toCurrency := context.GetString("toCurrency")

	if !exists {
		env.badRequest(context)
	}

	apiConverter := moneyconverter.CurrencyConverter{
		RateStrategy: moneyconverter.APIRateStrategy{},
		ToCurrency: model.Currency{
			CurrencyCode: toCurrency,
		},
	}

	moneyAdder := moneyadder.MoneyAdder{
		CurrencyConverter: apiConverter,
		Monies: monies.([]model.Money),
	}

	result, err := moneyAdder.CalculateSumResult()

	if err != nil {
		env.internalServerError(context, err)
	}

	response := network.Response {
		StatusCode: 200,
		Body: result,
	}

	env.respond(context, response)
}

func (env environment) repoSumMoney(context *gin.Context) {
	monies, exists := context.Get("monies")
	toCurrency := context.GetString("toCurrency")

	if !exists {
		env.badRequest(context)
	}

	repoConverter := moneyconverter.CurrencyConverter{
		RateStrategy: moneyconverter.RepositoryRateStrategy{Repo: env.repo},
		ToCurrency: model.Currency{
			CurrencyCode: toCurrency,
		},
	}

	moneyAdder := moneyadder.MoneyAdder{
		CurrencyConverter: repoConverter,
		Monies: monies.([]model.Money),
	}

	result, err := moneyAdder.CalculateSumResult()

	if err != nil {
		env.internalServerError(context, err)
	}

	response := network.Response {
		StatusCode: 200,
		Body: result,
	}

	env.respond(context, response)
}


func (env environment) respond(c *gin.Context, response network.Response) {
	obj := gin.H{}

	if response.Error != nil {
		obj["error"] = response.Error.Error()
		c.JSON(response.StatusCode, obj)
	}
	if response.Body != nil {
		c.JSON(response.StatusCode, response.Body)
	}
}

func (env environment) badRequest(context *gin.Context) {
	env.respond(context, network.Response{
		StatusCode: 400,
		Error:      errors.New("Bad request"),
	})

	context.Abort()
}

func (env environment) internalServerError(context *gin.Context, err error) {
	env.respond(context, network.Response{
		StatusCode: 500,
		Error:      err,
	})

	context.Abort()
}

//SumParams -
type SumParams struct {
	Monies []model.Money `json:"monies"`
	ToCurrency string
}
