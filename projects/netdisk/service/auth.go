package service

import (
	"netdisk/utils"
	"time"
)

type AuthHandler struct{}

var (
	AccessToken       string
	AccessTokenExpire time.Time
)

func (*AuthHandler) Login(wait *WaitConn, req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()

	if req.Username != Configuration.Username || req.Password != Configuration.Password {
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
