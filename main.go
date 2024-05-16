package main

import (
	"fmt"
	"http_on_brainfxxk/brainfxxk"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()

	e.Static("/", "public")

	e.GET("/" , GetHandler)
	e.POST("/run" , PostHandler)
	e.POST("/run-by-file", PostFileHandler)
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
	bf := brainfxxk.NewParser(req.Code, 128, 100000)
	if err := bf.Exec(); err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}
	return c.String(http.StatusOK, bf.OutputString)
}

func PostFileHandler(c echo.Context) error {
	file, err := c.FormFile("codefile")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	extension := filepath.Ext(file.Filename)
	extension = strings.TrimPrefix(extension, ".")
	if extension != "bf" {
		return c.String(http.StatusBadRequest, fmt.Errorf("unsupported file format").Error())
	}
	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	content, err := ioutil.ReadAll(src)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	bf := brainfxxk.NewParser(string(content), 128, 100000)
	if err := bf.Exec(); err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}
	return c.String(http.StatusOK, bf.OutputString)
}