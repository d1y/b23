// create by d1y<chenhonzhou@gmail.com>
// 公用方法

package path

import (
	"os"
	"os/user"
	"path"
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

// HomeDir 主目录
var HomeDir = ""

// DownloadDir 下载目录
var DownloadDir = "/Desktop"

func init() {
	HomeDir = GetHomePath()
	DownloadDir = path.Join(HomeDir, DownloadDir)
}
