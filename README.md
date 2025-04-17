# Go Gin Boilerplate ToDo

This is a Go Gin based API project to manage user controlled ToDo lists. The app itself is almost trivial and useless but, this project is designed to act as an opinionated boilerplate that can start and accelerate other backend web services using good architecture.

* Out of the box ready to use JWT security 
* Feature based code layout
* Easy to extend with new features

This project is built on top of the following great packages

* Application management through [Cobra](https://github.com/spf13/cobra), a package for creating powerful applications that start from the terminal
* [Bun](https://bun.uptrace.dev/) acting as a lightweight ORM (Object Relational Mapper) layer for data persistence including DB management and migration
* [Gin](https://gin-gonic.com/), a framework for building fast web services

## Prerequisites

* Go version greater than 1.12 in order to use modules

## Development Quickstart

There are sane development defaults provided by Cobra so, all you need to do is get the modules and run

* `go run main.go serve`

## Building

* Run `go build -o todoapi` to create the executable in project root

### Docker

There is a Dockerfile available for creating a production ready build of the Backend which includes only the built executable on top of [scratch](https://hub.docker.com/_/scratch) exposing port `8080`.

Take note that you will have to map the SQLite database in to the container yourself in order. Depending on where you map it, use the `-d, --sqlite-dsn string` option.

## Running in production

* Have `GIN_MODE=release` environment variable set
* Point the program to your SQLiteDB using the `-d, --sqlite-dsn string` program flags; see [Runtime configuration](#runtime-configuration) for more information 

## Runtime configuration

Runtime configuration is provided by Cobra so, for a list of options available, you can run `go run main.go help serve` which will display the following:

```shell
Flags:
      --auth-header-key string     Prefix to use with the Authorization header (default "Bearer")
      --auth-identity-key string   Identity key to use as part of JWT auth (default "id")
      --auth-key string            Secret key to use as salt as part of JWT auth (default "Secret Key")
      --auth-realm string          Security Realm to use with JWT (default "development zone")
      --auth-user-name string      Username to allow for authentication (default "admin")
      --auth-user-pass string      Password to allow for authentication (default "password")
  -h, --help                       help for serve
  -p, --port string                API port if not using env.PORT (default "8080")
  -d, --sqlite-dsn string          Data Source Name string for connecting SQLite DB (default "file:data.db?cache=shared")
```

## Extending this projecct

### Adding commands

The application is initiated using Cobra; there is also a CLI tool that will make adding new commands to your project simpler: https://github.com/spf13/cobra-cli
T
