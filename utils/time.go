package utils

import "time"



func TimeFormat(t int64, format string) string{
	return time.Unix(t,0).Format(format)
}


