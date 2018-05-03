package main

import (
	"net/http"
	"bufio"
	"fmt"
	"os"

	"github.com/labstack/echo"
)

/*type (
	Geolocation struct {
		Altitude  float64
		Latitude  float64
		Longitude float64
	}
)

var (
	locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{1899, 39.096849, -120.032351},
		{2619, 37.865101, -119.538329},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}
)*/

func main() {
	e := echo.New()
	file, err := os.Open("./citylots.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set(echo.HeaderContentEncoding, "chunked")
		c.Response().WriteHeader(http.StatusOK)

		for scanner.Scan() {
			c.Response().Write(scanner.Bytes())
			c.Response().Flush()
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":7000"))
}