package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"netdisk/utils"
	"path"
	"strings"
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

func (*FileHandler) Mkdir(wait *WaitConn, req struct {
	Path string `json:"path"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()

	if req.Path == "" {
		wait.SetResult("创建路径错误", nil)
		return
	}

	_, err := FilePtr.FileInfo.FindDir(req.Path, true)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}
}

func (*FileHandler) List(wait *WaitConn, req struct {
	Path string `json:"path"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()

	infos, err := FilePtr.FileInfo.FindDir(req.Path, false)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}
	items := make([]*item, 0, len(infos.FileInfos))
	for _, info := range infos.FileInfos {
		if info.IsDir || info.FileMD5 != "" {
			_item := &item{
				Filename: info.Name,
				IsDir:    info.IsDir,
				Date:     info.ModeTime,
			}
			if info.IsDir {
				_item.Size = "-"
			} else {
				_item.Size = utils.ConvertBytesString(info.FileSize)
			}
			items = append(items, _item)
		}
	}
	wait.SetResult("", &fileListData{
		DiskTotal:    FileDiskTotal,
		DiskTotalStr: utils.ConvertBytesString(FileDiskTotal),
		DiskUsed:     FilePtr.UsedDisk,
		DiskUsedStr:  utils.ConvertBytesString(FilePtr.UsedDisk),
		Total:        len(items),
		Items:        items,
	})
}

func (*FileHandler) Remove(wait *WaitConn, req struct {
	Path     string   `json:"path"`
	Filename []string `json:"filename"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()
	if req.Path == "" || len(req.Filename) == 0 {
		wait.SetResult("请求参数错误", nil)
		return
	}
	infos, err := FilePtr.FileInfo.FindDir(req.Path, false)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}

	for _, filename := range req.Filename {
		if err = Remove(infos, filename); err != nil {
			wait.SetResult(err.Error(), nil)
			return
		}
	}

	//todo: calUsedDisk()
	CalUsedDisk()
}

func (*FileHandler) Rename(wait *WaitConn, req struct {
	Path    string `json:"path"`
	OldName string `json:"oldName"`
	NewName string `json:"newName"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()
	if req.Path == "" || req.OldName == "" || req.NewName == "" {
		wait.SetResult("请求参数错误", nil)
		return
	}
	if req.OldName == req.NewName {
		return
	}
	if strings.Contains(req.NewName, "/") {
		wait.SetResult("文件名不能含有'/'", nil)
		return
	}
	dirInfo, err := FilePtr.FileInfo.FindDir(req.Path, false)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}

	srcInfo, ok := dirInfo.FileInfos[req.OldName]
	if !ok {
		wait.SetResult("文件不存在", nil)
		return
	}
	if err = Copy2src(srcInfo, dirInfo, req.NewName); err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}
	// 移除原文件
	_ = Remove(dirInfo, req.OldName)
	CalUsedDisk()
}

// 移动、复制文件或文件夹
func (*FileHandler) Mvcp(wait *WaitConn, req struct {
	Source []string `json:"source"`
	Target string   `json:"target"`
	Move   bool     `json:"move"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()
	if len(req.Source) == 0 || req.Target == "" {
		wait.SetResult("请求参数错误！", nil)
		return
	}
	tarDir, err := FilePtr.FileInfo.FindDir(req.Target, false)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}
	for _, source := range req.Source {
		srcPath, srcName := path.Split(source)
		srcDir, err := FilePtr.FileInfo.FindDir(srcPath, false)
		if err != nil {
			wait.SetResult(err.Error(), nil)
			return
		}
		srcInfo, ok := srcDir.FileInfos[srcName]
		if !ok {
			wait.SetResult("文件不存在", nil)
			return
		}
		// 不能移动到自身或子目录下
		if tarDir.AbsPath == srcDir.AbsPath || strings.Contains(tarDir.AbsPath, srcInfo.AbsPath) {
			wait.SetResult("不能拷贝、移动文件夹到自身目录或子目录", nil)
			return
		}
		if _, ok := tarDir.FileInfos[srcName]; ok {
			wait.SetResult("目标目录下已存在同名文件", nil)
			return
		}
		if err = Copy2src(srcInfo, tarDir, srcInfo.Name); err != nil {
			wait.SetResult(err.Error(), nil)
			return
		}
		if req.Move {
			// 移除原文件
			_ = Remove(srcDir, srcName)
		}
	}

	CalUsedDisk()
}

type DownloadArg struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

func Download(ctx *gin.Context, req *DownloadArg) {
	Logger.Infof("%s %v", GetCurrentRoute(ctx), req)
	info, err := FilePtr.FileInfo.FindDir(req.Path, false)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	file, ok := info.FileInfos[req.Filename]
	if !ok || file.FileMD5 == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	absPath := file.AbsPath
	if !SaveFileMultiple {
		// 虚拟保存，修正到真实文件路径
		md5File_, ok := FilePtr.MD5Files[file.FileMD5]
		if !ok {
			ctx.Status(http.StatusBadRequest)
			return
		}
		absPath = md5File_.File
	}
	// 设置响应的header头
	ctx.Writer.Header().Add("Content-type", "application/octet-stream")
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", req.Filename))
	ctx.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	ctx.File(absPath)
	ctx.Status(200)
}
