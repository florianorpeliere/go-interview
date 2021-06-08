package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/dailymotion/code-review-for-interviews/config"
	"github.com/dailymotion/code-review-for-interviews/database"
)

var (
	// flags
	printVersion bool

	// set by Makefile
	version   = "unknown"
	buildTime = "unknown"
)

func init() {
	flag.BoolVar(&printVersion, "version", false, "Print the version.")
}

func main() {
	flag.Parse()
	if printVersion {
		fmt.Printf("version %s build at %s (with %s)\n", version, buildTime, runtime.Version())
		os.Exit(0)
	}

	reloadInterval, _ := time.ParseDuration(config.Configuration.ReloadInterval)
	ticker := time.NewTicker(reloadInterval)

	go func() {
		for range ticker.C {
			currencies, err := loadCurrencies(config.Configuration.Currencies)
			if err != nil {
				log.Println(err)
			}
			for _, currency := range currencies {
				err = database.AddCurrency(currency)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	log.Println("starting http server", config.Configuration.ServerBindAddress)
	log.Fatal(http.ListenAndServe(config.Configuration.ServerBindAddress, http.HandlerFunc(handler)))
}

// http handler to get currencies
func handler(w http.ResponseWriter, r *http.Request) {
	currencies, ok := r.URL.Query()["currency"]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if len(currencies) != 1 {
		http.Error(w, "invalid value for the currency parameter", http.StatusBadRequest)
		return
	}

	currencyCode := currencies[0]
	currency, err := database.GetCurrency(currencyCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currency)
}
