# Go Restful API Boilerplate

[![GoDoc Badge]][godoc] [![GoReportCard Badge]][goreportcard]

Easily extendible RESTful API boilerplate aiming to follow idiomatic go and best practice.

The goal of this boiler is to have a solid and structured foundation to build upon on.

Any feedback and pull requests are welcome and highly appreciated. Feel free to open issues just for comments and discussions.

## Features

The following feature set is a minimal selection of typical Web API requirements:

- Configuration using [viper](https://github.com/spf13/viper)
- CLI features using [cobra](https://github.com/spf13/cobra)
- Structured logging with [zap](https://github.com/uber-go/zap)
- Routing with [Gin](https://github.com/gin-gonic/gin) and middleware
- Request data validation using [Go Validator](https://github.com/go-playground/validator)
- Redis support using [Go redis](https://github.com/go-redis/redis/)
- Enables request stats using [Gin stats](https://github.com/semihalev/gin-stats) on endpoint `/request/stats`
- Enables metrics and observability via [Go metrics](https://github.com/rcrowley/go-metrics)
- Makefile setup

## Start Application

- Clone this repository
- Run the application to see available commands: `make run`
- Run the application with command _serve_: `make server`