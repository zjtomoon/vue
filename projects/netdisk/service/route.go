package service

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/yddeng/utils/log"
	"github.com/yddeng/utils/task"
	"net/http"
	"netdisk/utils"
	"path"
	"reflect"
	"strings"
	"sync"
	"time"
)

var (
	app       *gin.Engine
	taskQueue *task.TaskPool
)

func Launch() {
	taskQueue = task.NewTaskPool(1, 1024)
	SaveFileMultiple = Configuration.SaveFileMultipe
	FileDiskTotal = Configuration.FileDiskTotal * utils.MB

	LoadFilePath(Configuration.FilePath)

	app = gin.New()
	app.Use(gin.Logger(), gin.Recovery())

	// 跨域
	app.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,PATCH")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Expose-Headers", "*")
		if ctx.Request.Method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	})

	if Configuration.StaticFS {
		app.StaticFS("/static", gin.Dir(Configuration.FilePath, true))
	}

	// 前端
	if Configuration.WebIndex != "" {
		app.Use(static.Serve("/", static.LocalFile(Configuration.WebIndex, false)))
		app.NoRoute(func(ctx *gin.Context) {
			ctx.File(Configuration.WebIndex + "/index.html")
		})
	}

	// todo: initHandler(app)
	initHandler(app)
	port := strings.Split(Configuration.WebAddr, ":")[1]
	webAddr := fmt.Sprintf("0.0.0.0:%s", port)

	Logger.Infof("start web service on %s", Configuration.WebAddr)

	if err := app.Run(webAddr); err != nil {
		log.Error(err)
	}
}

// 应答结构
type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WaitConn struct {
	code     int
	ctx      *gin.Context
	route    string
	result   Result
	done     chan struct{}
	doneOnce sync.Once
}

func newWaitConn(ctx *gin.Context, route string) *WaitConn {
	return &WaitConn{
		ctx:   ctx,
		code:  http.StatusOK,
		route: route,
		done:  make(chan struct{}),
	}
}

func (this *WaitConn) Done(code ...int) {
	this.doneOnce.Do(func() {
		if this.result.Message == "" {
			this.result.Success = true
		}
		if len(code) > 0 {
			this.code = code[0]
		}
		close(this.done)
	})
}

func (this *WaitConn) GetRoute() string {
	return this.route
}

func (this *WaitConn) Context() *gin.Context {
	return this.ctx
}

func (this *WaitConn) SetResult(message string, data interface{}) {
	this.result.Message = message
	this.result.Data = data
}

func (this *WaitConn) Wait() {
	<-this.done
}

type webTask func()

func (t webTask) Do() {
	t()
}

func transBegin(ctx *gin.Context, fn interface{}, args ...reflect.Value) {
	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		panic("value not func")
	}
	typ := val.Type()
	if typ.NumIn() != len(args)+1 {
		panic("func argument error")
	}
	route := GetCurrentRoute(ctx)
	wait := newWaitConn(ctx, route)
	if err := taskQueue.SubmitTask(webTask(func() {
		ok := checkToken(ctx, route)
		if !ok {
			wait.SetResult("Token验证失败", nil)
			wait.Done(401)
			return
		}
		val.Call(append([]reflect.Value{reflect.ValueOf(wait)}, args...))
	})); err != nil {
		wait.SetResult("访问人数过多", nil)
		wait.Done()
	}
	wait.Wait()
	ctx.JSON(wait.code, wait.result)
}

func GetCurrentRoute(ctx *gin.Context) string {
	return ctx.FullPath()
}

func getJsonBody(ctx *gin.Context, inType reflect.Type) (inValue reflect.Value, err error) {
	if inType.Kind() == reflect.Ptr {
		inValue = reflect.New(inType.Elem())
	} else {
		inValue = reflect.New(inType)
	}
	if err = ctx.ShouldBindJSON(inValue.Interface()); err != nil {
		return
	}
	if inType.Kind() != reflect.Ptr {
		inValue = inValue.Elem()
	}
	return
}

func WrapHandle(fn interface{}) gin.HandlerFunc {
	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		panic("value not func")
	}
	typ := val.Type()
	switch typ.NumIn() {
	case 1:
		return func(ctx *gin.Context) {
			transBegin(ctx, fn)
		}
	case 2:
		return func(ctx *gin.Context) {
			inValue, err := getJsonBody(ctx, typ.In(1))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Json unmarshal failed",
					"error":   err.Error(),
				})
				return
			}
			transBegin(ctx, fn, inValue)
		}
	default:
		panic("func symbol error")
	}
}

var (
	// 需要验证token的路由
	routeNeedToken = map[string]struct{}{
		"/file/list":    {},
		"/file/mkdir":   {},
		"file/remove":   {},
		"file/rename":   {},
		"file/mvcp":     {},
		"file/download": {},

		"/upload/check":  {},
		"/upload/upload": {},

		"/shared/create": {},
		"/shared/list":   {},
		"/shared/cancel": {},
	}
)

func checkToken(ctx *gin.Context, route string) bool {
	if _, ok := routeNeedToken[route]; !ok {
		return true
	}
	tkn := ctx.GetHeader("Access-Token")
	if tkn == "" {
		return false
	}
	if AccessTokenExpire.IsZero() || time.Now().After(AccessTokenExpire) {
		AccessToken = ""
		AccessTokenExpire = time.Time{}
		return false
	}

	return tkn == AccessToken
}

func initHandler(app *gin.Engine) {
	authHandle := new(AuthHandler)
	authGroup := app.Group("/auth")
	authGroup.POST("/login", WrapHandle(authHandle.Login))

	// todo: fileHandle
	fileHandle := new(FileHandler)
	fileGroup := app.Group("/file")
	fileGroup.POST("/mkdir", WrapHandle(fileHandle.Mkdir))
	fileGroup.POST("/list", WrapHandle(fileHandle.List))
	fileGroup.POST("/remove", WrapHandle(fileHandle.Remove))
	fileGroup.POST("/rename", WrapHandle(fileHandle.Rename))
	fileGroup.POST("/mvcp", WrapHandle(fileHandle.Mvcp))
	fileGroup.POST("/download", func(ctx *gin.Context) {
		var req *DownloadArg
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Json unmarshal failed",
				"error":   err.Error(),
			})
			return
		}
		Download(ctx, req)
	})
	// todo: uploadHandle
	uploadHandle := new(UploadHandler)
	uploadGroup := app.Group("/upload")
	uploadGroup.POST("/check", WrapHandle(uploadHandle.Check))
	uploadGroup.POST("/upload", WrapHandle(uploadHandle.Upload))

	// todo: shareHandle
	shareHandle := new(ShareHandler)
	shareGroup := app.Group("/shared")
	shareGroup.POST("/create", WrapHandle(shareHandle.Create))
	shareGroup.POST("/cancel", WrapHandle(shareHandle.Cancel))
	shareGroup.POST("/list", WrapHandle(shareHandle.List))
	shareGroup.POST("/info", WrapHandle(shareHandle.Info))
	shareGroup.POST("/path", WrapHandle(shareHandle.Path))
	shareGroup.POST("/download", func(ctx *gin.Context) {
		var req *ShareDownloadArg
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Json unmarshal failed",
				"error":   err.Error(),
			})
			return
		}
		shared, err := shareHandle.CheckShared(req.Key, req.SharedToken)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		if req.Path != shared.Path {
			children := false
			for _, name := range shared.Filename {
				if strings.Contains(req.Path, path.Join(shared.Path, name)) {
					children = true
					break
				}
			}
			if !children {
				ctx.Status(http.StatusBadRequest)
				return
			}
		}

		Download(ctx, &DownloadArg{
			Path:     req.Path,
			Filename: req.Filename,
		})
	})
}
