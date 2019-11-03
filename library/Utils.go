package library

import (
	"fmt"
	"net"
	"reflect"
	"regexp"
	"strings"
	"time"
)

//工具接口
type utils_i interface {
	//域名+端口转IP+端口
	Url2IPPort(url string, port string) ([]string, error)
	//根据正则 筛选数据
	Regexp(regex string,text string) ([]string)
	//时间格式化
	TimeFormat(dateTime time.Time) string
	//时间2时间戳
	Time2Unix(dateTime time.Time) int64
	//字符串->时间戳
	Str2Unix(formatTimeStr string) int64
	//字符串->时间对象
	Str2Time(formatTimeStr string) time.Time
	//时间对象->字符串
	Time2Str() string
	//时间对象->字符串
	Stamp2Str(stamp int64) string
	//时间戳->时间对象
	Stamp2Time(stamp int64)time.Time
	//list(slice) => string
	Implode(list interface{}, seq string) string
	//获取机器Ip4
	GetServerIp4() (ip string)
}
//工具类
type utils struct {

}
//工具类初始化
func newUtils() (u *utils) {
	return &utils{}
}
//域名+端口转IP+端口
func (u *utils) Url2IPPort(url string, port string) ([]string, error) {
	servers, err := net.LookupHost(url)
	if err != nil {
		return nil, err
	}
	ipports := make([]string, len(servers))
	for i, ip := range servers {
		ipports[i] = fmt.Sprintf("%s:%s", ip, port)
	}
	return ipports, nil
}

//根据正则 筛选数据
func (u *utils) Regexp(regex string,text string) ([]string) {
	//匹配话题
	text_regexp := regexp.MustCompile(regex)
	//wb_status.Text = "http://t.cn/EI6H1b7 快来参加微博社交电商，太棒了！！！"
	regexps := text_regexp.FindAllStringSubmatch(text,-1)
	regexps_len := len(regexps)
	arr := make([]string,0)
	if regexps_len > 0 { //正文中有匹配
		for i_reg := 0; i_reg < regexps_len; i_reg++ {
			regexps_len_sub := len(regexps[i_reg])
			for i_reg_sub := 1; i_reg_sub < regexps_len_sub; i_reg_sub++ {
				arr = append(arr, regexps[i_reg][i_reg_sub])
			}
		}
	}
	return arr;
}

/**
时间格式化
 */
func (u *utils) TimeFormat(dateTime time.Time) string {
	date := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), dateTime.Nanosecond(), time.Local)
	return date.Format("2006-01-02 15:04:05")
}

/**
 时间2时间戳
 */
func (u *utils) Time2Unix(dateTime time.Time) int64 {
	date := u.TimeFormat(dateTime)
	timeU, _ := time.Parse("2006-01-02 15:04:05", date)
	return timeU.Unix()
}

/**字符串->时间戳*/
func (u *utils) Str2Unix(formatTimeStr string) int64 {
	timeStruct := u.Str2Time(formatTimeStr)
	//毫秒
	//millisecond:=timeStruct.UnixNano()/1e6
	return  timeStruct.Unix()
}

/**字符串->时间对象*/
func (u *utils) Str2Time(formatTimeStr string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型
	return theTime

}

/**时间对象->字符串*/
func (u *utils) Time2Str() string {
	const shortForm = "2006-01-01 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

/*时间戳->字符串*/
func (u *utils) Stamp2Str(stamp int64) string {
	timeLayout := "2006-01-02 15:04:05"
	str := time.Unix(stamp/1000,0).Format(timeLayout)
	return str
}
/*时间戳->时间对象*/
func (u *utils) Stamp2Time(stamp int64)time.Time {
	stampStr := u.Stamp2Str(stamp)
	timer := u.Str2Time(stampStr)
	return timer
}

//list(slice) => string
func (u *utils) Implode(list interface{}, seq string) string {
	listValue := reflect.Indirect(reflect.ValueOf(list))
	if listValue.Kind() != reflect.Slice {
		return ""
	}
	count := listValue.Len()
	listStr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		v := listValue.Index(i)
		if str, err := getValue(v); err == nil {
			listStr = append(listStr, str)
		}
	}
	return strings.Join(listStr, seq)
}

func getValue(value reflect.Value) (res string, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		res, err = getValue(value.Elem())
	default:
		res = fmt.Sprint(value.Interface())
	}
	return
}

//获取机器Ip4
func (u *utils) GetServerIp4() (ip string)  {
	addrs,err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	ip = ""
	for _,value := range addrs {
		if ipnet,ok := value.(*net.IPNet);ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip;
}

