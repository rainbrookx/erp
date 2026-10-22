package util

import (
	"strconv"
	"strings"
	"time"
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

func ParseDateToStr(t *time.Time) string {
	if t == nil {
		return ""
	}
	// Go 的时间格式化模板固定为：2006-01-02 15:04:05
	return t.Format("2006-01-02")
}

func ParseDateToStr2(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func StringsToInt64s(ss []string) []int64 {
	result := make([]int64, 0, len(ss))
	for _, s := range ss {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}
		result = append(result, n)
	}
	return result
}
