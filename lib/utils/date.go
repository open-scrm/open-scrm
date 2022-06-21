package utils

import "time"

const (
	ymdhis = "2006-01-02 15:04:05"
)

func GetNow() string {
	return time.Now().Format(ymdhis)
}

func FromUnix(n int64) string {
	return time.Unix(n, 0).Format(ymdhis)
}
