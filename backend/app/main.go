package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/h4shu/shiritori-go/db"
	"github.com/h4shu/shiritori-go/handler"
	"github.com/h4shu/shiritori-go/service"
)

func main() {
	port := ":" + os.Getenv("PORT")

	rdb := db.NewRedisClient()
	defer rdb.Close()

	wsvc := service.NewWordchainService(rdb)
	sh := handler.NewWordchainHandler(wsvc)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/wc", sh.GetWordchain)
	e.POST("/wc", sh.AddWordchain)
	e.Logger.Fatal(e.Start(port))
}
