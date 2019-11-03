package model

import (
	"wb-impulse-sender/conf"
	"wb-impulse-sender/library"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"sync"
	_ "github.com/go-sql-driver/mysql"
)


var daoMap sync.Map

//初始化
func newInstance(cfg *conf.Config, driver string) (instance *gorm.DB, err error) {
	var ins *gorm.DB
	value,ok := daoMap.Load(driver)
	if ok {
		switch t := value.(type) {
		case *gorm.DB:
			ins = value.(*gorm.DB)
		default:
			_ = t
		}
	}else {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",cfg.Mysql[driver].User,cfg.Mysql[driver].Pwd,cfg.Mysql[driver].Host,cfg.Mysql[driver].Port,cfg.Mysql[driver].Db)
		ins, err = gorm.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		ins.DB().SetMaxIdleConns(10)
		ins.DB().SetMaxOpenConns(1000)
		ins.SingularTable(true)

		//开启日志
		ins.LogMode(true)
		daoMap.Store(driver, ins)
	}
	return ins, nil
}

//获取查询条件
func getQueryWhere(conditions map[string]interface{}) (string, []interface{}) {
	cols := make([]string, 0, len(conditions))
	values := make([]interface{}, 0, len(conditions))

	for col,value := range conditions {
		query := strings.Join([]string{col, " = ?"},"")
		cols = append(cols, query)
		values = append(values,value)
	}
	query := library.GetLibraryInstance().Utils.Implode(cols," and ")

	return query, values
}