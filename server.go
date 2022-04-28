package main

import (
	"fmt"
	"net/http"
	"Go_App/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/data", func(c echo.Context) error {
		router.StoreCovidData()
		return c.String(http.StatusOK, "Data loaded Successfully!")
	})

	e.GET("/data/location", func(c echo.Context) error {
		
		latitude := c.QueryParam("lat")
		longitude := c.QueryParam("long")
		fmt.Println("lat: ", latitude, "long: ", longitude)
		s := string(router.RetrieveDataFromLatLong(latitude,longitude))
		return c.String(http.StatusOK, s)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
