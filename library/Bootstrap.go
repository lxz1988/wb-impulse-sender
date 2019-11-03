package library

import "sync"

var once sync.Once

var instance library

//library结构体
type library struct {
	Exception exception_i
	Logger logger_i
	Utils utils_i
}

/**
	获取library实例
 */
func GetLibraryInstance() (*library) {
	once.Do(func() {
		instance.Exception = newException()
		instance.Logger = newLogger()
		instance.Utils = newUtils()

	})
	return &instance
}
