package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

// 本来可以试着 echo.Echo.Get("/",Decorator(something))来传递东西的，但是似乎没必要这么麻烦。
//func Decorator(something string, f echo.HandlerFunc) echo.HandlerFunc {
//	fmt.Println(s)
//	return f
//}

func PATCH(c echo.Context) error {
	return nil
}

func POST(c echo.Context) error {
	return nil
}

func DELETE(c echo.Context) error {
	return nil
}
func GET(c echo.Context) error {
	return c.String(200, "test")
}
func PUT(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

//留作外部调用接口
func LoadRouters() {
	e.GET(confuse+"/:id", GET)
	e.POST(confuse+"/", POST)
	e.PUT(confuse+"/", PUT)
	e.PATCH(confuse+"/", PATCH)
	e.DELETE(confuse+"/:id", DELETE)
}

func StartEchoHandle() {
	// Echo instance

	e = echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	LoadRouters()

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
