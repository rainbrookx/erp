package util

import (
	"strconv"
	"strings"
)

// StrToLongList "1,2,3,4,5,6" -> []int64{1,2,3,4,5,6}
func StrToLongList(strArr string) []int64 {
	idList := make([]int64, 0)
	if strArr == "" {
		return idList
	}
	for _, s := range strings.Split(strArr, ",") {
		s = strings.TrimSpace(s) // 顺带去掉两端空白，比原 Java 版更稳
		if s == "" {
			continue
		}
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue // 解析失败的项跳过；如要严格报错可改成返回 error
		}
		idList = append(idList, n)
	}
	return idList
}
