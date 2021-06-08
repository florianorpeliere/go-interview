# Readme

* REST API -> GET /currency/{code}

# Main

* parseDuration err not checked
* loading loop: err logged but flow continue
* http server no timeouts configured
* handler errors not consistent (404, 400)
* internal err msg exposed to the outside world

## currency

* load in // ? goroutines ? errors handling ?
* http client no timeouts configured

# Config

* global var
* init, flag.parse
* if name == "" -> nothing
* yaml err unchecked

# Database

* global var
* init
* panic

# Dependencies

* no vendor, no dep tool

# Tests

* no tests

# Docker

* entrypoint missing config
* run as root by default

# Follow-up questions

* code organisation ?
* ready for prod ? something missing ? (metrics for example)
