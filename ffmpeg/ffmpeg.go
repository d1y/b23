// create by d1y<chenhonzhou@gmail.com>
// go bind ffmpeg wrapper

package ffmpeg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	goPath "path"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/d1y/b23/fastgit"
	"github.com/d1y/b23/path"
	"github.com/d1y/b23/utils"
	"github.com/imroc/req"
)

// ConvertFormat2mp3 将 视频格式转为 `mp3`
func ConvertFormat2mp3(enter string, output string) bool {
	bin, err := FindFfmpegBin()
	if err != nil {
		panic("找不到 ffmpeg")
	}
	var outputFile = fmt.Sprintf("%v.mp3", output)
	var execRun = exec.Command(bin, "-i", enter, outputFile)
	err = execRun.Run()
	if err != nil {
		// panic("执行错误, 写入的文件可能已经存在!!")
		return false
	}
	return true
}

// FindFfmpegBin 返回 `ffmpeg` 的路径
func FindFfmpegBin() (string, error) {
	var runtimeIsFfmpeg = utils.IsCommandAvailable(`ffmpeg`)
	if runtimeIsFfmpeg {
		return `ffmpeg`, nil
	}
	var cBind = path.GetLocalWrapperFfmpegPath()
	if cBind != "" {
		return cBind, nil
	}
	return "", errors.New("命令不存在")
}

// DownloadFfmpegBin 下载 `ffmpeg` 二进制文件
// linux/osx 平台不推荐使用!!只适配 `windows`, 因为 `unix` 安装很简单..
func DownloadFfmpegBin() error {
	var ctxPath = path.GetCurrentPath()
	f, err := ioutil.TempFile("", "b23")
	if err != nil {
		panic("create temp file is error")
	}
	defer os.Remove(f.Name())
	var Fpath = f.Name()
	var targetPath = goPath.Join(ctxPath, "lib/")
	s := spinner.New(spinner.CharSets[33], 100*time.Millisecond)
	s.Start()
	s.Suffix = "下载中."
	var url = fastgit.CreateFastGitURL(`/vot/ffbinaries-prebuilt/releases/download/v4.2.1/ffmpeg-4.2.1-win-64.zip`)
	progress := func(current, total int64) {
		var ctx = fmt.Sprintf("%v%v", float32(current)/float32(total)*100, "%")
		s.Prefix = ctx
	}
	r, err := req.Get(url, req.DownloadProgress(progress))
	if err != nil {
		return errors.New("download ffmpeg bin file is error")
	}
	fmt.Println("download file path: ", Fpath)
	r.ToFile(Fpath)
	s.Stop()
	e := utils.Unzip(Fpath, targetPath)
	if e != nil {
		return errors.New("unzip ffmpeg bin file is error")
	}
	return nil
}

// windows 下载 `ffmpeg`
func windowsInstallFfmpeg() error {
	if runtime.GOOS != "windows" {
		return nil
	}
	_, err := FindFfmpegBin()
	if err != nil {
		return DownloadFfmpegBin()
	}
	return nil
}

func init() {
	windowsInstallFfmpeg()
}
