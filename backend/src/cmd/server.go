package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/h4shu/shiritori-go/adapters/controllers"
	"github.com/h4shu/shiritori-go/adapters/gateways/apis"
	"github.com/h4shu/shiritori-go/adapters/gateways/repositories"
	"github.com/h4shu/shiritori-go/adapters/presenters"
	"github.com/h4shu/shiritori-go/application/intractors"
	"github.com/h4shu/shiritori-go/domain/entities"
	"github.com/h4shu/shiritori-go/infrastructure/db/redis"
	"github.com/h4shu/shiritori-go/infrastructure/web/clients/groq"
	"github.com/h4shu/shiritori-go/infrastructure/web/handlers"
)

func Server(port string) {

	rdb := redis.NewRedisClient()
	defer rdb.Close()

	g := groq.NewGroqClient()
	a := apis.NewValidateWordApi(g)
	wt := entities.WordTypeHiragana
	ws := redis.NewWordchainStore(rdb, "wordchain")
	wr := repositories.NewWordchainRepository(ws, wt)
	wu := intractors.NewWordchainUsecase(a, wr, wt, 1000)
	wp := presenters.NewWordchainPresenter()
	wc := controllers.NewWordchainController(wu, wp, wt)
	sh := handlers.NewWordchainHandler(wc)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/wc", sh.List)
	e.GET("/wc/last", sh.GetLast)
	e.POST("/wc", sh.Append)
	e.Logger.Fatal(e.Start(":" + port))
}
