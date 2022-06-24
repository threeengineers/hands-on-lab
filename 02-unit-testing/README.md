# Hands-on Lab / Unit Testing

In this lab you will create a simple calculator API using Go programming language.

## Pre-requisite

Please make sure that you have the latest version of Golang. We recommend to use Go v1.18.3.

## Getting Started

Before you are getting started, please clone this repo.

```
git clone git@github.com:threeengineers/hands-on-lab.git
```

### Open project

Go to the project directory.

```
cd hands-on-lab/02-unit-testing
```

### Install dependencies

In this project, we are using external library. Now, install the deps.

```
go mod vendor
```

### Start the app

Now that you have installed all deps, let's start the app.

```
go run main.go
```

You should see the following log.

```

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.7.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

### Access the API

There are two available APIs, as following.

#### Calculator - Add

```
curl http://localhost:1323/api/v1/calculator/add?a=1&b=1
```

#### Calculator - Subtract

```
curl http://localhost:1323/api/v1/calculator/subtract?a=5&b=1
```

## Author

[Faris](faris@monolog.id)