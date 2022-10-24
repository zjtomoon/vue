package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"myBlog/controllers"
)

func InitRouter(app *iris.Application)  {
	baseUrl := "/api"
	mvc.New(app.Party(baseUrl+"/user")).Handle(controllers.NewUserController())
	mvc.New(app.Party(baseUrl+"/category")).Handle(controllers.NewCategoryController())
	mvc.New(app.Party(baseUrl+"/article")).Handle(controllers.NewArticleController())
	mvc.New(app.Party(baseUrl+"/auth")).Handle(controllers.NewAuthController())
	mvc.New(app.Party(baseUrl+"/comment")).Handle(controllers.NewCommentController())

	//app.Use(middleware.GetJWT().Serve)
}
