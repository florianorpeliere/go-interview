# Currencies Store

**Currencies Store** is a [Go](http://golang.org/) web-application that retrieves some currencies rates from an external provider, stores them in a [PostgreSQL](https://www.postgresql.org/) database, and exposes an HTTP REST API to get the latest rate for a currency.

It retrieves data from <http://free.currencyconverterapi.com/> for the configured currencies, at a regular interval (defined in the configuration).

### Usage / Configuration

See `config/config.yaml` for a configuration sample. When you start the app, use the `-config` parameter to pass the path of the configuration file:

```
currencies -config config/config.yaml
```

### REST API

The REST API to retrieve a currency is very simple: just do a `GET` request on `/`, and pass the `currency` query string. The value should be the currency code (`EUR`, `GBP`, ...) to retrieve.

```
curl "http://localhost:8080/?currency=EUR"
```

The result is a JSON object:

```
{
    "name": "EUR",
    "rate": 1.067,
    "date": "2017-04-05T14:00:27.014902Z"
}
```

### Building

* `make build` to build the binary in `output/currencies`
* `make build-in-docker` to build inside a docker container (if you don't have go installed)
* `make build-docker-image` to build a docker image (will build the binary first)

### Database

* create the database with `createdb currencies`
* create the table with `psql currencies < database/creation.sql`

### Dev setup

* run `go get ./...` in the project directoy to retrieve all the dependencies
* run `go build` to build the binary
