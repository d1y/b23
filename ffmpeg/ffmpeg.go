// create by d1y<chenhonzhou@gmail.com>
// go bind ffmpeg wrapper

package ffmpeg

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/d1y/b23/path"
	"github.com/d1y/b23/utils"
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
