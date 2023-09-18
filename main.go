package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var currencies = []string{"JPY", "USD", "TWD"}

type currencyConversion struct {
	Currencies map[string]map[string]float64 `json:"Currencies"`
}

var conversion currencyConversion

type response struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

func main() {
	c := gin.Default()
	c.GET("/", CurrencyConversionHandler)
	c.Run()
}

func CurrencyConversionHandler(c *gin.Context) {
	var res response
	currenciesJsonStr := `{
		"currencies": {
			"TWD": {
				"TWD": 1,
				"JPY": 3.669,
				"USD": 0.03281
			},
			"JPY": {
				"TWD": 0.26956,
				"JPY": 1,
				"USD": 0.00885
			},
			"USD": {
				"TWD": 30.444,
				"JPY": 111.801,
				"USD": 1
			}
		}
	}`

	if err := json.Unmarshal([]byte(currenciesJsonStr), &conversion); err != nil {
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	currencyMap := make(map[string]map[string]float64)
	for source, val := range conversion.Currencies {
		currencyMap[source] = make(map[string]float64)
		for target, rate := range val {
			currencyMap[source][target] = rate
		}
	}

	sourceCurrency := c.DefaultQuery("source", "")
	targetCurrency := c.DefaultQuery("target", "")
	amountStr := c.DefaultQuery("amount", "")

	if !strings.Contains(amountStr, "$") {
		res.Msg = "Amount Format Error"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	amountStr = strings.ReplaceAll(amountStr, ",", "")
	amountStr = strings.ReplaceAll(amountStr, "$", "")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		res.Msg = "Amount Format Error"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if !isCurrencyVal(sourceCurrency) {
		res.Msg = "Source Format Error"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if !isCurrencyVal(targetCurrency) {
		res.Msg = "Target Format Error"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	formattedAmount := fmt.Sprintf("%.2f", currencyMap[sourceCurrency][targetCurrency]*amount)
	parts := strings.Split(formattedAmount, ".")
	integerPart := parts[0]
	decimalPart := parts[1]
	formattedAmount = "$" + addCommas(integerPart) + "." + decimalPart
	res.Msg = "success"
	res.Amount = formattedAmount
	c.JSON(http.StatusOK, res)
}

func addCommas(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return addCommas(s[:n-3]) + "," + s[n-3:]
}

func isCurrencyVal(currency string) bool {
	for _, c := range currencies {
		if c == currency {
			return true
		}
	}
	return false
}
