# Go server
---

* [About project](#about-project)
* [Install](#install)
* [Using](#using)

---

## About project

Project consists of the following packages:
* handlers
* config
* connection
* entity
* main

The handler package consists of router handlers using gorilla-mux.
This packet is responsible for all the routing logic.

The config package is used to load and use .env files for server configuration.

The connection package is used for creation session between server and database and running database migrations.

The entity package is used for database interaction. All CRUD requests go through this package.

The main package is used for configure and starting server

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/AA55hex/golang_bootcamp
```
## Using

There is Makefile for using project. Use `make` for help

### Launch

Use `make go-run` to run the docker-compose in attachment mode.
There you can see server-related log information (after conteiners initialization)

### Testing

Use `make go-test` to test project packages
Tests are available for router handlers, connections, configuration packages.

### Clearing

Use `make clear` to down docker-compose conteiners
