package main

import (
	"encoding/json"
	"fmt"
	"go-test/controllers/fines"
	"go-test/models"
	"go-test/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Test struct {
	Name int `json:"name"`
}

func main() {

	err := models.Connect()
	if err != nil {
		panic(err)
	}

	// array json
	eTest := Test{}
	var arr []interface{}
	arr = append(arr, map[string]interface{}{"name": "soup", "team": "Eieiza"})
	arr = append(arr, map[string]interface{}{"name": "max", "team": "Eieiza"})
	for _, value := range arr {
		toBtye, err := json.Marshal(value)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(toBtye, &eTest)
		fmt.Println(eTest.Name)
	}

	// declare variable // if else
	name := ""
	var number int
	var pNumber *int
	if name == "" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	if number == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	if pNumber == nil {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// services
	var a *int
	var b *int
	var c int
	var d int
	c = 1
	d = 2
	a = &c
	b = &d
	sum := services.Summary(a, b)
	fmt.Println(sum)

	e := echo.New()
	e.Debug = true
	e.Use(middleware.CORS())

	// FinesRoutes(e.Group("/fines"))

	e.Use(Middleware)
	{
		e.POST("/add", fines.Add)
		// api.POST("/getAll", fines.GetAll)
		// api.POST("/getArray", fines.GetArray)
	}

	e.Logger.Fatal(e.Start(":8080"))
}

// func ApiRoutes(api *echo.Group) {
// 	FinesRoutes(api.Group("/fines"))
// }

// func FinesRoutes(finesRoutes *echo.Group) {
// 	finesRoutes.GET("/add", fines.Add)
// }

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		invalid := false

		if invalid == false {
			return next(c)
		} else {
			byte, err := json.Marshal(map[string]interface{}{"code": 400, "status": "bad request", "value": nil})
			if err != nil {
				fmt.Println(err)
			}
			return c.JSONBlob(http.StatusOK, byte)
		}
	}
}
