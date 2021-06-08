package database

import (
	"github.com/dailymotion/code-review-for-interviews/model"
	"github.com/pkg/errors"
	"gopkg.in/mgutz/dat.v2/dat"
)

// GetCurrency retrieves the latest currency rate from the database
func GetCurrency(currencyCode string) (*model.Currency, error) {
	currency := &model.Currency{}

	err := DB.
		Select(columnName, columnPriceInDollars, columnDate).
		From(tableName).
		Where(dat.Eq{columnName: currencyCode}).
		OrderBy(columnDate + " DESC").
		Limit(1).
		QueryStruct(currency)
	if err != nil {
		return nil, errors.Wrapf(err, "could not query database for currency %s", currencyCode)
	}

	return currency, nil
}
