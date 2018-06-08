package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/galenchurch/gobone"
)

type LedState struct {
	LED   string `json:"led"`
	State string `json:"state"`
}

func onOff(led gobone.LED, c string) {
	if c == "on" {
		led.LedOn()
	} else if c == "off" {
		led.LedOff()
	} else {
		log.Printf("%s - not a valid state for an led\n", led.Sys.Name)
	}
}

func LedHandler(cxt echo.Context) error {

	l := new(LedState)

	l.LED = cxt.Param("led")

	if err := cxt.Bind(l); err != nil {
		return cxt.JSON(http.StatusBadRequest, `{ "statusCode": 400, "error": "Bad data!", "message": "unknown state!" }`)
	}

	switch l.LED {
	case gobone.Usr0.Sys.Name:
		onOff(gobone.Usr0, l.State)
	case gobone.Usr1.Sys.Name:
		onOff(gobone.Usr1, l.State)
	case gobone.Usr2.Sys.Name:
		onOff(gobone.Usr2, l.State)
	case gobone.Usr3.Sys.Name:
		onOff(gobone.Usr3, l.State)
	}

	return cxt.JSON(http.StatusOK, l)
}
