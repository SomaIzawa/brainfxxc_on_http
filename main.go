package main

import (
	"http_on_brainfxxk/brainfxxk"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()

	e.Static("/", "public")

	e.GET("/" , GetHandler)
	e.POST("/run" , PostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func GetHandler(c echo.Context) error {
	return c.File("public/index.html")
}

type PostReq struct {
	Code string `json:"code"`
}

func PostHandler(c echo.Context) error {
	var req PostReq
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	// brainfxxk
	bf := brainfxxk.NewParser(req.Code, 128, 10000)
	if err := bf.Exec(); err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}
	return c.String(http.StatusOK, bf.OutputString)
}