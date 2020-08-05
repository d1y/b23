package main

import (
	"fmt"
	"os"

	"github.com/d1y/b23/utils"

	"github.com/d1y/b23/api"
	"github.com/d1y/b23/cli"
	"github.com/fatih/color"
)

func easy(id string) {
	if utils.IsValidURL(id) {
		ctx, idErr := utils.GetB23ID(id)
		if idErr != nil {
			panic(idErr)
		}
		id = ctx
	}
	info, err := api.GetB23VideoPagelist(id)
	if err != nil {
		color.Cyan("获取视频分p失败")
		os.Exit(1)
	}
	videoInfo, err := api.GetB23VideoInfo(id)
	if err != nil {
		color.Red("获取视频信息失败")
		os.Exit(1)
	}
	var title = videoInfo.Data.Title
	var firstCid = info.Data[0].Cid
	var cid = fmt.Sprintf("%v", firstCid)
	var urls = api.EasyGetB23VideoURL(id, cid)
	if len(urls) >= 1 {
		api.DownloadFileAndToMp3(urls[0], title)
	}
}

func main() {
	var args = os.Args
	if len(args) < 2 {
		color.Red(cli.Help())
		os.Exit(0)
	} else {
		easy(args[1])
	}
}
