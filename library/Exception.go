package library

import "fmt"

//异常接口
type exception_i interface {
	Error(er interface{}) (err error)
}
//异常结构体
type exception struct {

}
//初始化异常类
func newException() (ex *exception) {
	return &exception{}
}
//异常error方法
func (ex *exception) Error(er interface{}) (err error)  {
	err = fmt.Errorf("%v",er)
	//TODO 记录日志
	return err
}
