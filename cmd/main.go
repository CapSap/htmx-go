package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

}