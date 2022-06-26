package server

import "github.com/labstack/echo/v4"


func initRouter()*echo.Echo{
e := echo.New()
e.POST("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}