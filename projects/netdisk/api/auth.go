package api

import (
	"netdisk/conf"
	"netdisk/route"
	"netdisk/service"
	"netdisk/utils"
	"time"
)

type AuthHandler struct{}

var (
	AccessToken       string
	AccessTokenExpire time.Time
)

func (*AuthHandler) Login(wait *route.WaitConn, req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}) {
	service.Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()

	if req.Username != conf.Configuration.Username || req.Password != conf.Configuration.Password {
		wait.SetResult("用户或密码错误", nil)
		return
	}

	now := time.Now()
	if AccessToken == "" || now.After(AccessTokenExpire) {
		AccessToken = utils.GenToken(20)
		AccessTokenExpire = now.Add(time.Hour * 8)
	}

	wait.SetResult("", struct {
		Token string `json:"token"`
	}{
		Token: AccessToken,
	})
}
