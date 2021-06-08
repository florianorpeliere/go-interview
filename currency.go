package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"github.com/dailymotion/code-review-for-interviews/model"
)

func loadCurrencies(currencyCodes []string) ([]model.Currency, error) {
	currencies := make([]model.Currency, len(currencyCodes))
	for i, currencyCode := range currencyCodes {
		currency, err := loadCurrency(currencyCode)
		if err != nil {
			return nil, err
		}
		currencies[i] = *currency
	}
	return currencies, nil
}

func loadCurrency(currencyCode string) (*model.Currency, error) {
	query := fmt.Sprintf("%s_USD", currencyCode)
	url := fmt.Sprintf("http://free.currencyconverterapi.com/api/v3/convert?q=%s&compact=ultra", query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "could not query currencyconverterapi.com for currency %s", currencyCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("invalid http response status code for currency %s: %d (%s)", currencyCode, resp.StatusCode, resp.Status)
	}

	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "could not parse currencyconverterapi.com response for currency %s", currencyCode)
	}

	currency := model.Currency{
		Name:           currencyCode,
		PriceInDollars: gjson.ParseBytes(json).Get(query).Float(),
		Date:           time.Now(),
	}
	return &currency, nil
}
