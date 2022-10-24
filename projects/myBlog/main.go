package main

import (
	"flag"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"myBlog/conf"
	"myBlog/route"
)

func main() {
	flag.Parse()
	app := newApp()
	route.InitRouter(app)
	err := app.Run(iris.Addr(":"+conf.Sysconfig.Port),iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // allow everything,use that to change the hosts
		AllowCredentials: true,
		AllowedHeaders: []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)
	return app
}
