package config

import "fmt"

// base api
var baseAPI = `http://api.bilibili.com/`

// VideoInfo 视频详细信息
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/info.md
var VideoInfo = fmt.Sprintf("%v%v", baseAPI, "x/web-interface/view")

// VideoPagelist 视频分`p`
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/info.md
var VideoPagelist = fmt.Sprintf("%v%v", baseAPI, "x/player/pagelist")

// VideoPlayURL 获取视频流URL
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/videostream_url.md
var VideoPlayURL = fmt.Sprintf("%v%v", baseAPI, "x/player/playurl")

// B23Hosts b站镜像域名站点
var B23Hosts = []string{
	"bilibili.com",
	"www.bilibili.com",
	"b23.tv",
}
