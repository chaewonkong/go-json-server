package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/chaewonkong/go-json-server/utils"
	"github.com/labstack/echo/v4"
)

func readDir() (files []os.FileInfo, err error) {
	files, err = ioutil.ReadDir("./db")
	return
}

func main() {

	s, err := utils.ReadDirAndGetRouteConfig("./db")

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	for _, config := range s {
		config.AddRoute(e)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
