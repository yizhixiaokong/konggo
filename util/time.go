package util

import (
	"fmt"
	"konggo/pkg/convert"
	"time"
)

// 定义时间格式
const TimeFormat string = "2006-01-02 15:04:05.000"
const TimeNormalFormat string = "2006-01-02 15:04:05"
const ReportTimeFormat string = "20060102150405"
const ReportDateFormat string = "20060102"

// 获取时间
func GetTime() string {
	return time.Now().Format(TimeFormat)
}

func GetDate() string {
	now := time.Now()
	year := now.Year()
	month := now.Month() //time.Now().Month().String()
	day := now.Day()
	return fmt.Sprintf("%04d%02d%02d", year, month, day)
}

// 时间字符串转time.Time
func GetTimeFromString(ts string) (t time.Time, err error) {
	temp := []byte(ts + "+08:00")
	temp[10] = 'T'
	t, err = time.Parse(time.RFC3339, string(temp))
	return
}

// 日期字符串转time.Time
func GetDateFromString(date string) (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339, string(date+"T00:00:00+08:00"))
	return
}

// 获取一个远古时间,便于时间字段传空时插入数据库
func GetDefaultTime() (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339, string("1601-01-01T00:00:00+08:00"))
	return
}

func GetReportTime() string {
	return GetReportTimeBy(time.Now())
}

func GetReportDate() string {
	return GetReportDateBy(time.Now())
}

func GetReportTimeBy(time time.Time) string {
	return time.Format(ReportTimeFormat)
}

func GetReportDateBy(time time.Time) string {
	return time.Format(ReportDateFormat)
}

// 时间加n天
func AddTime(t time.Time, n int) time.Time {
	dur := time.Hour * time.Duration(24*n)
	return t.Add(dur)
}

// 获取当天0点的时间
func GetToday() (today time.Time) {
	now := time.Now()
	today = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return today
}

func GetNilTime() time.Time {
	return convert.ToTime(-1)
}
