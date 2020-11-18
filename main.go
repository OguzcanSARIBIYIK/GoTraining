package main

import (
	"Echo1/user"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello World!")
	})
	e.GET("/name/:name", func(context echo.Context) error {
		name := context.Param("name")

		return context.String(200, name)
	})
	e.GET("/query", func(context echo.Context) error {
		name := context.QueryParam("name")

		return context.String(http.StatusOK, name)
	})
	e.POST("/form", func(context echo.Context) error {
		name := context.FormValue("name")
		surname := context.FormValue("surname")

		return context.JSON(http.StatusOK, map[string]string{
			"name":    name,
			"surname": surname,
		})
	})
	a := e.Group("/admin")
	//a.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Skipper: func(c echo.Context) bool {
	//		if c.QueryParam("token") == "99" {
	//			return true
	//		}
	//		return false
	//	},
	//}))

	a.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {

		return func(context echo.Context) error {
			if context.QueryParam("token") != "99" {
				return context.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Yetkisiz işlem...",
				})

			}

			return handlerFunc(context)
		}

	})

	a.GET("/user/info", user.Info)
	a.GET("/user/detail/:id", user.Detail)
	a.POST("/user/delete", user.Delete)
	e.POST("/form/user", user.Create)
	e.Logger.Fatal(e.Start(":1323"))
}
