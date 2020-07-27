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