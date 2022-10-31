package api

import (
	"netdisk/route"
	"netdisk/service"
)

type fileListData struct {
	DiskUsed     uint64  `json:"diskUsed"`
	DiskUsedStr  string  `json:"diskUsedStr"`
	DiskTotal    uint64  `json:"diskTotal"`
	DiskTotalStr string  `json:"diskTotalStr"`
	Total        int     `json:"total"`
	Items        []*item `json:"items"`
}

type item struct {
	Filename string `json:"filename"`
	IsDir    bool   `json:"isDir"`
	Size     string `json:"size"`
	Date     string `json:"date"`
}

type FileHandler struct{}

func (*FileHandler) Mkdir(wait *route.WaitConn, req struct {
	Path string `json:"path"`
}) {
	service.Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()

	if req.Path == "" {
		wait.SetResult("创建路径错误", nil)
		return
	}

	_, err := service.FilePtr.FileInfo.FindDir(req.Path, true)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}
}

func (*FileHandler) List(wait *route.WaitConn, req struct {
	Path string `json:"path"`
}) {
	service.Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()

	info, err := service.FilePtr.FileInfo.FindDir(req.Path, false)
}
