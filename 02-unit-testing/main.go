package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/threeengineers/hands-on-lab/02-unit-testing/feature"
)

func main() {
	e := echo.New()

	e.GET("/api/v1/hello", func(c echo.Context) error {
		var result string = "hello world"

		if os.Getenv("HELLO_MSG") != "" {
			result = os.Getenv("HELLO_MSG")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"result": result})
	})

	e.GET("/api/v1/calculator/add", func(c echo.Context) error {
		var err error

		a, err := strconv.ParseInt(c.QueryParam("a"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		b, err := strconv.ParseInt(c.QueryParam("b"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		calculator := new(feature.Calculator)
		result := calculator.Add(int32(a), int32(b))

		return c.JSON(http.StatusOK, map[string]interface{}{"result": result})
	})

	e.GET("/api/v1/calculator/subtract", func(c echo.Context) error {
		var err error

		a, err := strconv.ParseInt(c.QueryParam("a"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		b, err := strconv.ParseInt(c.QueryParam("b"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		calculator := new(feature.Calculator)
		result := calculator.Subtract(int32(a), int32(b))

		return c.JSON(http.StatusOK, map[string]interface{}{"result": result})
	})

	var port string
	if os.Getenv("APP_PORT") == "" {
		port = ":1323"
	} else {
		port = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	}
	e.Logger.Fatal(e.Start(port))
}
