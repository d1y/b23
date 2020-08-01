package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/d1y/b23/config"
	"github.com/d1y/b23/ffmpeg"
	"github.com/d1y/b23/path"
	"github.com/d1y/b23/utils"
	cmdColor "github.com/fatih/color"
	"github.com/imroc/req"
)

// createMiddlewareHeader 创建 `header`
func createMiddlewareHeader(vid string) req.Header {
	referer := fmt.Sprintf("https://www.bilibili.com/video/%v", vid)
	if vid == "" {
		referer = "https://www.bilibili.com"
	}
	hd := req.Header{
		"referer":    referer,
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
	}
	return hd
}

// genIDKey 生成不同 `id` 类型的 `key`
func genIDKey(vid string) string {
	var key = "bvid"
	if utils.CheckBID(vid) {
		if utils.CheckIDIsAvid(vid) {
			key = "avid"
		}
	}
	return key
}

// GetB23VideoPagelist 获取视频多`p`(主要是用来获取`cid`)
func GetB23VideoPagelist(vid string) (BilibiliVideoPagelistJSON, error) {
	var key = genIDKey(vid)
	qs := req.Param{
		key: vid,
	}
	hd := createMiddlewareHeader(vid)
	resp, err := req.Get(config.VideoPagelist, qs, hd)
	var respBody BilibiliVideoPagelistJSON
	if err != nil {
		return respBody, errors.New("获取请求失败")
	}
	var r = resp.Response().Body
	body, err := ioutil.ReadAll(r)
	json.Unmarshal(body, &respBody)
	return respBody, nil
}

// GetB23VideoInfo 获取视频信息
func GetB23VideoInfo(id string) (BilibiliVideoInfoJSON, error) {
	var key = genIDKey(id)
	qs := req.Param{
		key: id,
	}
	hd := createMiddlewareHeader(id)
	resp, err := req.Get(config.VideoInfo, qs, hd)
	var respBody BilibiliVideoInfoJSON
	if err != nil {
		return respBody, errors.New("获取请求失败")
	}
	var r = resp.Response().Body
	body, err := ioutil.ReadAll(r)
	json.Unmarshal(body, &respBody)
	return respBody, nil
}

// GetB23VideoURL 获取视频链接 `flv`
func GetB23VideoURL(vid string, cid string) (BilibiliVideoURLJSON, error) {
	var key = genIDKey(vid)
	hd := createMiddlewareHeader(vid)
	qs := req.Param{
		key:   vid,
		"cid": cid,
	}
	resp, err := req.Get(config.VideoPlayURL, hd, qs)
	var respBody BilibiliVideoURLJSON
	if err != nil {
		return respBody, errors.New("获取请求失败")
	}
	var r = resp.Response().Body
	body, err := ioutil.ReadAll(r)
	json.Unmarshal(body, &respBody)
	return respBody, nil
}

// EasyGetB23VideoURL 获取视频链接 `flv`, 直接返回 `[]string`
func EasyGetB23VideoURL(vid string, cid string) []string {
	ctx, err := GetB23VideoURL(vid, cid)
	if err != nil {
		return []string{}
	}
	var lists = ctx.Data.Durl
	// utils.ContentLength2MB(lists.Size)
	if len(lists) < 1 {
		return []string{}
	}
	var list = lists[0]
	var result = []string{}
	result = append(result, list.URL)
	var backURL = list.BackupURL
	for i := 0; i < len(backURL); i++ {
		result = append(result, backURL[i])
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// DownloadFileAndToMp3 下载文件然后转为 `mp3`
func DownloadFileAndToMp3(url string, output string) {
	f, err := ioutil.TempFile("", "b23")
	if err != nil {
		panic("create temp file is error")
	}
	defer os.Remove(f.Name())
	var Fpath = f.Name()
	var tempFilePrint = fmt.Sprintf("临时路径: %v", Fpath)
	cmdColor.Cyan(tempFilePrint)
	check(err)
	s := spinner.New(spinner.CharSets[33], 100*time.Millisecond)
	s.Start()
	s.Suffix = "下载中."
	progress := func(current, total int64) {
		var ctx = fmt.Sprintf("%v%v", float32(current)/float32(total)*100, "%")
		s.Prefix = ctx
	}
	hd := createMiddlewareHeader("")
	r, _ := req.Get(url, hd, req.DownloadProgress(progress))
	r.ToFile(Fpath)
	s.Stop()
	cmdColor.Yellow("下载完成")
	s.Start()
	s.Prefix = "正在转换格式"
	var outputFilePath = path.CreateDesktopFile(output)
	var outputFlag = ffmpeg.ConvertFormat2mp3(Fpath, outputFilePath)
	if !outputFlag {
		panic("转换失败, 可能是写入的文件可能已经存在!!")
	}
	s.Stop()
	var resultMessage = fmt.Sprintf("转换成功: %v", outputFilePath)
	cmdColor.Green(resultMessage)
}
