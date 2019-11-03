package library

import "fmt"

//日志接口
type logger_i interface {
	Log(content interface{})
}
//日志结构体
type logger struct {

}
//初始化日志类
func newLogger() (*logger) {
	return &logger{}
}
//log方法
func (l *logger) Log(content interface{}) {

	fmt.Println(content)
}
