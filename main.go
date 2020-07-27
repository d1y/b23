package main

import (
	"fmt"
	"os"

	"github.com/d1y/b23/api"
	"github.com/fatih/color"
)

func main() {
	var args = os.Args
	if len(args) < 2 {
		color.Red("请传递参数")
		os.Exit(0)
	} else {
		var id = args[1]
		info, err := api.GetB23VideoPagelist(id)
		if err != nil {
			color.Cyan("请求错误")
		}
		var firstCid = info.Data[0].Cid
		var cid = fmt.Sprintf("%v", firstCid)
		var urls = api.EasyGetB23VideoURL(id, cid)
		if len(urls) >= 1 {
			api.DownloadFileAndToMp3(urls[0])
		}
	}
}
