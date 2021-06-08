package database

import (
	"database/sql"
	"time"

	"github.com/dailymotion/code-review-for-interviews/config"
	_ "github.com/lib/pq"
	"gopkg.in/mgutz/dat.v2/dat"
	"gopkg.in/mgutz/dat.v2/sqlx-runner"
)

const (
	tableName            = "currencies"
	columnName           = "name"
	columnPriceInDollars = "dollars_rate"
	columnDate           = "created_at"
)

// global database (pooling provided by SQL driver)
var DB *runner.DB

func init() {
	// create a normal database connection through database/sql
	db, err := sql.Open("postgres", config.Configuration.DatabaseURI)
	if err != nil {
		panic(err)
	}

	// ensures the database can be pinged with an exponential backoff (15 min)
	runner.MustPing(db)

	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	// set this to enable interpolation
	dat.EnableInterpolation = true

	// set to check things like sessions closing.
	// Should be disabled in production/release builds.
	dat.Strict = false

	// Log any query over 10ms as warnings. (optional)
	runner.LogQueriesThreshold = 10 * time.Millisecond

	DB = runner.NewDB(db, "postgres")
}
