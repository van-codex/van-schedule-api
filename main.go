package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	_ "net/http/pprof"
	"schedule-api/elastic"
	"schedule-api/router"
	"schedule-api/task"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	es := elastic.Init(cfg.Section("elasticsearch"))
	routes := router.Init(
		task.Inject(es),
		es,
	)
	app.Post("put", routes.PutRoute)
	app.Post("get", routes.GetRoute)
	app.Post("lists", routes.ListsRoute)
	app.Post("all", routes.AllRoute)
	app.Post("running", routes.RunningRoute)
	app.Post("delete", routes.DeleteRoute)
	app.Run(iris.Addr(":3000"), iris.WithoutServerError(iris.ErrServerClosed))
}