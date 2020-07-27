// create by d1y<chenhonzhou@gmail.com>
// 公用方法

package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// CheckBID 判断是不是 `bilibili` id
func CheckBID(bid string) bool {
	var worA = "av"
	var worB = "bv"
	var isA = strings.Index(bid, worA) > -1
	var isB = strings.Index(bid, worB) > -1
	return (isA || isB)
}

// CheckIDIsAvid 判断是不是 `avid`
func CheckIDIsAvid(id string) bool {
	var worx = "av"
	var index = strings.Index(id, worx)
	return index > -1
}

// ContentLength2MB 将 `ContentLength` 转为 `xxmb` 格式
func ContentLength2MB(n int64) string {
	var mb = n / 1024 / 1024
	f, err := strconv.ParseFloat(string(mb), 2)
	if err != nil {
		return "未知"
	}
	var r = fmt.Sprintf("%vmb", f)
	return r
}
