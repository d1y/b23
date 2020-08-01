package api

// BilibiliVideoPagelistJSON 视频分`p` json
type BilibiliVideoPagelistJSON struct {
	Code int64 `json:"code"`
	Data []struct {
		Cid       int64 `json:"cid"`
		Dimension struct {
			Height int64 `json:"height"`
			Rotate int64 `json:"rotate"`
			Width  int64 `json:"width"`
		} `json:"dimension"`
		Duration int64  `json:"duration"`
		From     string `json:"from"`
		Page     int64  `json:"page"`
		Part     string `json:"part"`
		Vid      string `json:"vid"`
		Weblink  string `json:"weblink"`
	} `json:"data"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
}

// BilibiliVideoURLJSON 视频 `url` json
type BilibiliVideoURLJSON struct {
	Code int64 `json:"code"`
	Data struct {
		AcceptDescription []string `json:"accept_description"`
		AcceptFormat      string   `json:"accept_format"`
		AcceptQuality     []int64  `json:"accept_quality"`
		Durl              []struct {
			Ahead     string   `json:"ahead"`
			BackupURL []string `json:"backup_url"`
			Length    int64    `json:"length"`
			Order     int64    `json:"order"`
			Size      int64    `json:"size"`
			URL       string   `json:"url"`
			Vhead     string   `json:"vhead"`
		} `json:"durl"`
		Format       string `json:"format"`
		From         string `json:"from"`
		Message      string `json:"message"`
		Quality      int64  `json:"quality"`
		Result       string `json:"result"`
		SeekParam    string `json:"seek_param"`
		SeekType     string `json:"seek_type"`
		Timelength   int64  `json:"timelength"`
		VideoCodecid int64  `json:"video_codecid"`
	} `json:"data"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
}

// BilibiliVideoInfoJSON `bilibili` 视频信息
type BilibiliVideoInfoJSON struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		State     int    `json:"state"`
		Attribute int    `json:"attribute"`
		Duration  int    `json:"duration"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		NoCache bool `json:"no_cache"`
		Pages   []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
	} `json:"data"`
}
