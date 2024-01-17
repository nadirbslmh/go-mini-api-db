package main

import (
	"go-mini-api-db/database"
	"go-mini-api-db/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
