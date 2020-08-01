// create by d1y<chenhonzhou@gmail.com>
// 公用方法

package path

import (
	"os"
	"os/user"
	"path"

	"github.com/d1y/b23/utils"
)

// GetHomePath 获取用户主目录
func GetHomePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		currentUser, _ := user.Current()
		home = currentUser.HomeDir
	}
	return home
}

// CreateDesktopFile 创建桌面文件
func CreateDesktopFile(filename string) string {
	var x = path.Join(DownloadDir, filename)
	return x
}

// GetCurrentPath 获取当前路径..
func GetCurrentPath() string {
	curr, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return curr
}

// GetLocalWrapperFfmpegPath 获取懒人包的 `ffmpeg`
// 主要会在 `windows` 平台打包一个懒人包, 自带了 `ffmpeg`
func GetLocalWrapperFfmpegPath() string {
	var curr = GetCurrentPath()
	var ffmpeg = path.Join(curr, "./lib/ffmpeg/bin/ffmpeg.exe")
	var fffmpegExist = utils.CheckFileIsExists(ffmpeg)
	if fffmpegExist {
		return ffmpeg
	}
	return ""
}

// HomeDir 主目录
var HomeDir = ""

// DownloadDir 下载目录
var DownloadDir = "/Desktop"

func init() {
	HomeDir = GetHomePath()
	DownloadDir = path.Join(HomeDir, DownloadDir)
}
