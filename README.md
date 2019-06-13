# money-converter

Hello!

This program is a service that let you sum different exchanges.

## To run it:

1) Create .env file with the "PORT" field that you are going to run the server inside the /cmd directory.
2) cd cmd
3) go run main.go


## How to use it:

### Get exchanges

  API:
  GET http://host:port/exchanges/api

  REPO:
  GET http://host:port/exchanges/repo

### Convert money

  API:
  GET http://host:port/exchanges/api/convert?from=USD&to=JPY&amount=10

  REPO:
  GET http://host:port/exchanges/repo/convert?from=USD&to=ARS&amount=10
 
### Sum money

  API: 
  POST http://host:port/exchanges/api/sum
 
```json
{
	"monies":
	[
			{
				"amount":10,
				"currencyCode":"USD"
			},
			{
				"amount":130,
				"currencyCode":"NOK"
			},
			{
				"amount":10,
				"currencyCode":"BGN"
			}
	],
	"toCurrency": "USD"
}
```
  REPO: 
  POST http://host:port/exchanges/repo/sum
 
```json
{
	"monies":
	[
			{
				"amount":10,
				"currencyCode":"USD"
			},
			{
				"amount":130,
				"currencyCode":"ARS"
			},
			{
				"amount":10,
				"currencyCode":"ARS"
			}
	],
	"toCurrency": "USD"
}
```
