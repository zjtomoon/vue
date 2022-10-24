package controllers

import (
	"github.com/kataras/iris/v12"
	"log"
	"myBlog/models"
	"myBlog/service"
)

type AuthController struct {
	Ctx iris.Context
	Service service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		Service: service.NewAuthService(),
	}
}

func (g *AuthController) PostUserInfo() (result models.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:",err)
	}
	return g.Service.GetUserInfo(m)
}