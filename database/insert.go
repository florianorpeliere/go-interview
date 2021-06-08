package database

import (
	"github.com/dailymotion/code-review-for-interviews/model"
)

// AddCurrency inserts the currency data into the database
func AddCurrency(currency model.Currency) error {
	_, err := DB.
		InsertInto(tableName).
		Pair(columnName, currency.Name).
		Pair(columnPriceInDollars, currency.PriceInDollars).
		Pair(columnDate, currency.Date).
		Exec()
	return err
}
