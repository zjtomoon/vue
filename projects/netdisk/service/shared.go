package service

import (
	"errors"
	"netdisk/utils"
	"path"
	"strings"
	"time"
)

type fileShare struct {
	Key         string   `json:"key"` // 路由
	Route       string   `json:"route"`
	Path        string   `json:"path"`
	Filename    []string `json:"filename"`    // 分享的文件、文件夹
	SharedToken string   `json:"sharedToken"` //访问密码
	CreateTime  int64    `json:"createTime"`  // 分享时间，单位秒
	Deadline    int64    `json:"deadline"`    // 到期时间戳，单位秒
	Looked      int      `json:"looked"`      // 浏览次数
}

type ShareDownloadArg struct {
	Key         string `json:"key"`
	SharedToken string `json:"sharedToken"`
	Path        string `json:"path"`
	Filename    string `json:"filename"`
}

// 服务器重启后，清空
var fileShared = map[string]*fileShare{}

type ShareHandler struct{}

func (*ShareHandler) GetSharedRoute(shared *fileShare) string {
	return "http://" + Configuration.WebAddr + "/shared/s/" + shared.Key
}

func (this *ShareHandler) List(wait *WaitConn) {
	defer func() {
		wait.Done()
	}()
	wait.SetResult("", fileShared)
}

func (this *ShareHandler) Create(wait *WaitConn, req struct {
	Path     string   `json:"path"`
	Filename []string `json:"filename"`
	Deadline int      `json:"deadline"` // 分享时间，单位天。0表示永久
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()
	if req.Path == "" || len(req.Filename) == 0 {
		wait.SetResult("请求参数错误", nil)
		return
	}
	dirInfo, err := FilePtr.FileInfo.FindDir(req.Path, false)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}
	for _, filename := range req.Filename {
		_, ok := dirInfo.FileInfos[filename]
		if !ok {
			wait.SetResult("文件（夹）不存在", nil)
			return
		}
	}

	key := utils.GenToken(16)
	for {
		if _, ok := fileShared[key]; !ok {
			break
		} else {
			key = utils.GenToken(16)
		}
	}
	now := time.Now()
	shared := &fileShare{
		Key:         key,
		Path:        req.Path,
		Filename:    req.Filename,
		SharedToken: utils.GenToken(4),
		CreateTime:  now.Unix(),
		Deadline:    0,
	}
	shared.Route = this.GetSharedRoute(shared)
	if req.Deadline > 0 {
		shared.Deadline = now.Add(time.Hour * 24 * time.Duration(req.Deadline)).Unix()
	}
	fileShared[key] = shared

	wait.SetResult("", struct {
		Route       string `json:"route"`
		SharedToken string `json:"sharedToken"`
		Deadline    int64  `json:"deadline"`
	}{
		Route:       shared.Route,
		SharedToken: shared.SharedToken,
		Deadline:    shared.Deadline,
	})
}

func (this *ShareHandler) Cancel(wait *WaitConn, req struct {
	Keys []string `json:"keys"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()
	for _, key := range req.Keys {
		delete(fileShared, key)
	}
}

func (this *ShareHandler) CheckShared(key, token string) (*fileShare, error) {
	shared, ok := fileShared[key]
	if !ok || (shared.Deadline != 0 && time.Now().Unix() > shared.Deadline) {
		delete(fileShared, key)
		return nil, errors.New("分享链接已过期")
	}

	if shared.SharedToken != token {
		return nil, errors.New("提取码错误")
	}
	_, err := FilePtr.FileInfo.FindDir(shared.Path, false)
	if err != nil {
		delete(fileShared, key)
		return nil, errors.New("分享链接已取消")
	}
	return shared, nil
}

func (this *ShareHandler) Info(wait *WaitConn, req struct {
	Key         string `json:"key"`
	SharedToken string `json:"sharedToken"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)
	defer func() {
		wait.Done()
	}()
	shared, err := this.CheckShared(req.Key, req.SharedToken)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}
	shared.Looked++

	type resp struct {
		Root       string  `json:"root"`
		Name       string  `json:"name"`
		CreateTime int64   `json:"createTime"`
		Deadline   int64   `json:"deadline"`
		Items      []*item `json:"items"`
	}
	ret := &resp{
		Root:       shared.Path,
		Name:       shared.Filename[0],
		CreateTime: shared.CreateTime,
		Deadline:   shared.Deadline,
	}
	if len(shared.Filename) > 1 {
		ret.Name += "等"
	}
	dirInfo, err := FilePtr.FileInfo.FindDir(shared.Path, false)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}
	items := make([]*item, 0, len(shared.Filename))
	for _, name := range shared.Filename {
		info, ok := dirInfo.FileInfos[name]
		if ok && (info.IsDir || info.FileMD5 != "") {
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

	ret.Items = items
	wait.SetResult("", ret)
}

func (this *ShareHandler) Path(wait *WaitConn, req struct {
	Key         string `json:"key"`
	SharedToken string `json:"sharedToken"`
	Path        string `json:"path"`
}) {
	Logger.Infof("%s %v", wait.GetRoute(), req)

	defer func() {
		wait.Done()
	}()

	shared, err := this.CheckShared(req.Key, req.SharedToken)
	if err != nil {
		wait.SetResult(err.Error(), nil)
		return
	}

	if req.Path == shared.Path {
		// 根目录
		dirInfo, err := FilePtr.FileInfo.FindDir(req.Path, false)
		if err != nil {
			wait.SetResult(err.Error(), nil)
			return
		}
		items := make([]*item, 0, len(shared.Filename))
		for _, name := range shared.Filename {
			info, ok := dirInfo.FileInfos[name]
			if ok && (info.IsDir || info.FileMD5 != "") {
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
		wait.SetResult("", struct {
			Items []*item `json:"items"`
		}{Items: items})
	} else {
		// 子目录
		children := false
		for _, name := range shared.Filename {
			if strings.Contains(req.Path, path.Join(shared.Path, name)) {
				children = true
				break
			}
		}
		if !children {
			wait.SetResult("路径不存在", nil)
			return
		}
		dirInfo, err := FilePtr.FileInfo.FindDir(req.Path, false)
		if err != nil {
			wait.SetResult(err.Error(), nil)
			return
		}
		items := make([]*item, 0, len(dirInfo.FileInfos))
		for _, info := range dirInfo.FileInfos {
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

		wait.SetResult("", struct {
			Items []*item `json:"items"`
		}{Items: items})
	}

}
