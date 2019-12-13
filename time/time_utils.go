package utils

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"strconv"
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
	case "MM-dd hh:mm":
		timeStr = time.Now().Format("15时04分05秒")
	default:
		timeStr = time.Now().String()
	}

	return timeStr
}

// Time2Str 时间转字符串
func Time2Str(unixTime int64, pattern string) string {
	t := time.Unix(unixTime, 0)

	var timeStr string

	switch pattern {
	case "yyyy-MM-dd":
		timeStr = t.Format("2006-01-02")
	case "yyyy/MM/dd":
		timeStr = t.Format("2006/01/02")
	case "yyyy-MM-dd hh:mm:ss":
		timeStr = t.Format("2006-01-02 15:04:05")
	case "hh:mm:ss":
		timeStr = t.Format("15:04:05")
	case "yyyy/MM/dd hh:mm:ss":
		timeStr = t.Format("2006/01/02 15:04:05")
	case "yyyyMMddhhmmss":
		timeStr = t.Format("20060102150405")
	case "年月日":
		timeStr = t.Format("2006年01月02日")
	case "年/月/日":
		timeStr = t.Format("2006年/01月/02日")
	case "年月日 时分秒":
		timeStr = t.Format("2006年01月02日 15时04分05秒")
	case "时分秒":
		timeStr = t.Format("15时04分05秒")
	case "yy-MM-dd hh:mm":
		timeStr = t.Format("06-01-02 15:04")
	case "MM-dd hh:mm":
		timeStr = t.Format("01-02 15:04")
	default:
		timeStr = t.String()
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

/**
* @des 时间转换函数
* @param atime string 要转换的时间戳（秒）
* @return string
 */
func StrTime(unixTime int64) string {
	var byTime = []int64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"年前", "天前", "小时前", "分钟前", "秒钟前"}
	now := time.Now().Unix()
	ct := now - unixTime
	if ct <= 0 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = mergeString(tempStr, unit[i]) //此处调用了一个我自己封装的字符串拼接的函数（你也可以自己实现）
		}
		break //我想要的形式是精确到最大单位，即："2天前"这种形式，如果想要"2天12小时36分钟48秒前"这种形式，把此处break去掉，然后把字符串拼接调整下即可（别问我怎么调整，这如果都不会我也是无语）
	}
	return res
}

/**
* @des 拼接字符串
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func mergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}
