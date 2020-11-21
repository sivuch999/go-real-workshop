package fines

import (
	"encoding/json"
	"fmt"
	"go-test/entities"
	"go-test/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) (err error) {
	eFines := entities.Fines{}
	var response interface{}

	err = c.Bind(&eFines)
	if err != nil {
		log.Println(err)
	}

	err = models.GormDB.Debug().Create(&eFines).Error
	if err != nil {
		fmt.Println(err)
		response = map[string]interface{}{"code": 400, "status": "bad request", "value": nil}
	} else {
		response = map[string]interface{}{"code": 200, "status": "ok", "value": eFines}
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSONBlob(http.StatusOK, bytes)
	// return nil
}
