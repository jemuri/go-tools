package utils

import (
	"context"
	"fmt"
	"time"

	"net"
)

//ConvertTime2Unix time string convert to unix
func ConvertTime2Unix(timeStr string) int64 {
	ttt, _ := time.Parse("2006-01-02", timeStr)
	return ttt.Unix()
}

//NowTimeStr 返回当前时间字符串
func NowTimeStr(format string) string {
	var timeStr string

	switch format {
	case "yyyy-MM-dd":
		timeStr = time.Now().Format("2006-01-02")
	case "yyyy/MM/dd":
		timeStr = time.Now().Format("2006/01/02")
	case "yyyy-MM-dd hh:mm:ss":
		timeStr = time.Now().Format("2006-01-02 15:04:05")
	case "hh:mm:ss":
		timeStr = time.Now().Format("15:04:05")
	case "yyyy/MM/dd hh:mm:ss":
		timeStr = time.Now().Format("2006/01/02 15:04:05")
	case "yyyyMMddhhmmss":
		timeStr = time.Now().Format("20060102150405")
	case "年月日":
		timeStr = time.Now().Format("2006年01月02日")
	case "年/月/日":
		timeStr = time.Now().Format("2006年/01月/02日")
	case "年月日 时分秒":
		timeStr = time.Now().Format("2006年01月02日 15时04分05秒")
	case "时分秒":
		timeStr = time.Now().Format("15时04分05秒")
	default:
		timeStr = time.Now().String()
	}

	return timeStr
}

//ComputeSub compute the time difference about the spend in function
func ComputeSub(ctx context.Context, startTime time.Time, funcName string) {
	duration := time.Since(startTime)
	fmt.Printf("LocalIP :%s,function: %s,spend time: %s", getLocalIP(), funcName, duration)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "InterfaceAddrs error !"
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
