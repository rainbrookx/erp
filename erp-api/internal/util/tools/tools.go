package tools

import "time"

const (
	DayFirstTime = " 00:00:00"
	DayLastTime  = " 23:59:59"
)

// Tools 工具方法，等价原Java Tools
type Tools struct{}

// GetNow 获取当前日期字符串，格式参考：yyyy-MM-dd
func (t *Tools) GetNow() string {
	return time.Now().Format("2006-01-02")
}

// GetYesterday 获取昨天日期 yyyy-MM-dd
func (t *Tools) GetYesterday() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
}

// GetCurrentMonth 获取本月，这里返回当月第一天 yyyy-MM-dd，适配 firstDayOfMonth
func (t *Tools) GetCurrentMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// FirstDayOfMonth 获取当月第一天 yyyy-MM-dd
func (t *Tools) FirstDayOfMonth(monthTime time.Time) string {
	return time.Date(monthTime.Year(), monthTime.Month(), 1, 0, 0, 0, 0, monthTime.Location()).Format("2006-01-02")
}

// GetYearBegin 获取本年第一天 yyyy-MM-dd
func (t *Tools) GetYearBegin() string {
	now := time.Now()
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
}

// GetYearEnd 获取本年最后一天 yyyy-MM-dd
func (t *Tools) GetYearEnd() string {
	now := time.Now()
	return time.Date(now.Year(), 12, 31, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
}

// StrToDate 将日期的字符串格式转为时间格式 "yyyy-MM-dd HH:mm:ss"
func StrToDate(dateString string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, dateString)
}
