package utils

import "time"

// StringTime 输入指定格式出生日期转换成time类型进行存储
func StringTime(birthday string) time.Time {
	t, _ := time.Parse("2006-01-02", birthday)
	return t
}

// TimeString 从数据存储的出生日期time类型转成字符串
func TimeString(t time.Time) string {
	return t.Format("2006-01-02")
}
