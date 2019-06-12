package ratescalc

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
	"errors"
)

type envelope struct {
	Cube []struct {
		Date  string `xml:"time,attr"`
		Rates []struct {
			Currency string `xml:"currency,attr"`
			Rate     string `xml:"rate,attr"`
		} `xml:"Cube"`
	} `xml:"Cube>Cube"`
}

//GetExchangeRate return currency rate from European Central Bank API. 
//The rates are based on EUR.
func GetExchangeRate(currencyName string) (float64, error) {
	var exchange float64

	env, err := getRatesEnvelope()

	if err != nil {
		return -1, err
	}

	founded := false

	for _, v := range env.Cube[0].Rates {
		if v.Currency == currencyName && !founded {
			exchange, err = strconv.ParseFloat(v.Rate, 64)
			founded = true
		}
	}

	if err != nil {
		return -1, err
	}

	if !founded {
		return -1, errors.New("Unexisting currency name")
	}

	return exchange, nil
}

//GetEnabledExchanges -
func GetEnabledExchanges() ([]string, error) {
	var exchanges []string

	env, err := getRatesEnvelope()

	if err != nil {
		return exchanges, err
	}

	for _, v := range env.Cube[0].Rates {
		exchanges = append(exchanges, v.Currency)
	}

	return exchanges, nil
}

func getRatesEnvelope() (envelope, error) {
	var (
		err error
		env envelope
	)

	resp, err := http.Get("http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		return env, err
	}

	defer resp.Body.Close()

	xmlCurrenciesData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return env, err
	}

	err = xml.Unmarshal(xmlCurrenciesData, &env)

	if err != nil {
		return env, err
	}

	return env, nil
}