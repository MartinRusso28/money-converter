package moneyconverter

import (
	"money-converter/pkg/converter/model"
	r "money-converter/internal/repository"
	"github.com/gin-gonic/gin"
	"strconv"
	"errors"
	"money-converter/internal/network"
	"money-converter/pkg/rates"
)

//GetMainEngine return money converter's server.
func GetMainEngine(repo r.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	env := environment{repo: repo}

	router.GET("/exchanges/api/convert",env.setConvertParams, env.apiConvertExchange)
	router.GET("/exchanges/api/", env.getAPIEnabledExchanges)
	
	router.GET("/exchanges/repo/convert",env.setConvertParams, env.repoConvertExchange)
	router.GET("/exchanges/repo/", env.getRepoEnabledExchanges)
	router.POST("/exchanges/sum/", env.getMonies, env.sumMoney)

	return router
}

type environment struct {
	repo r.Repository
}

func (env environment) getAPIEnabledExchanges(context *gin.Context){
	enabledEx, err := ratescalc.GetEnabledExchanges()

	if err != nil {
		env.internalServerError(context, err)
	}

	env.respondEnabledExchanges(context, enabledEx)
}

func (env environment) getRepoEnabledExchanges(context *gin.Context){
	enabledEx := env.repo.Money.GetEnabledExchanges()
	env.respondEnabledExchanges(context, enabledEx)
}

func (env environment) respondEnabledExchanges(context *gin.Context, exchanges []string){
	response := network.Response{
		StatusCode: 200,
		Body: exchanges,
	}

	env.respond(context, response)
}

func (env environment) setConvertParams(context *gin.Context) {
	fromCurrency := context.Query("from")
	toCurrency := context.Query("to")
	moneyAmount := context.Query("amount")

	moneyAmountFormatted, err := strconv.Atoi(moneyAmount)

	if err != nil || fromCurrency == "" || toCurrency == "" || moneyAmountFormatted <= 0{
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
		Currency:model.Currency{
			CurrencyCode: context.GetString("from"),
		},
	}

	converter := CurrencyConverter{
		RateStrategy: APIRateStrategy{},
		ToCurrency : toCurrency,
	}

	convertedMoney, err := converter.Convert(&fromMoney)

	if err != nil {
		env.internalServerError(context, err)
	}

	response := network.Response{
		StatusCode: 200,
		Body: convertedMoney,
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
		Currency:model.Currency{
			CurrencyCode: context.GetString("from"),
		},
	}

	converter := CurrencyConverter{
		RateStrategy: RepositoryRateStrategy{env.repo},
		ToCurrency : toCurrency,
	}

	convertedMoney, err := converter.Convert(&fromMoney)

	if err != nil {
		env.internalServerError(context, err)
	}

	response := network.Response{
		StatusCode: 200,
		Body: convertedMoney,
	}

	env.respond(context, response)
}


func (env environment) getMonies(context *gin.Context) {
	moniesRequest := moniesRequest{}
	context.BindJSON(&moniesRequest)

	for _,v := range moniesRequest.monies {
		if !v.Valid() {
			env.badRequest(context)
			return
		}
	}

	context.Set("monies", moniesRequest.monies)
}

func (env environment) sumMoney(context *gin.Context) {

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

type moniesRequest struct {
	monies []model.Money
}