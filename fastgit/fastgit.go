package fastgit

import (
	"fmt"
	"strings"
)

// CreateFastGitURL 创建一个 `hub.fastgit.org` 镜像站的 `alias-url`
func CreateFastGitURL(path string) string {
	var isX = strings.Contains(path, "/")
	if !isX {
		path = fmt.Sprintf("/%v", path)
	}
	var ctx = fmt.Sprintf("https://hub.fastgit.org%v", path)
	return ctx
}
