package main

import (
	"fmt"
	"waysbeans/database"
	"waysbeans/pkg/mysql"
	"waysbeans/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DataBaseInit()
	database.RunMigration()

	routes.Routes(e.Group("api/v1"))

	port := "5000"
	fmt.Println("server running on port", port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}
