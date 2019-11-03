package model

import (
	"github.com/jinzhu/gorm"
	"wb-impulse-sender/conf"
	"time"
)


type WorkNode struct {
	db *gorm.DB
	node
}

//数据库实体对象
type node struct {
	ID int64 `gorm:"column:id;PRIMARY_KEY"`
	ServerId string `gorm:"column:server_id"`
	Path string `gorm:"column:path"`
	Extra string `gorm:"column:extra"`

	WorkNodeCtime time.Time `gorm:"column:work_node_ctime"`
	WorkNodeMtime time.Time `gorm:"column:work_node_mtime"`
}
//指定数据库对应的表名
func (node) TableName() string {
	return "pond_work_node"
}

//初始化
func NewWorkNode(cfg *conf.Config, driver string) (wn *WorkNode)  {
	db, _ := newInstance(cfg, driver)
	wn = &WorkNode{
		db : db,
	}
	return wn
}

//查询单条
func (wn *WorkNode) Find() (*node) {
	f := node{}
	wn.db.Where("").Find(&f)
	return &f
}

//查询
func (wn *WorkNode) Select(conditions map[string]interface{}) ([]*node) {
	feeds := make([]*node, 0)
	query, where := getQueryWhere(conditions)
	wn.db.Where(query, where...).Find(&feeds)
	return feeds
}

//创建
func (wn *WorkNode) Insert() {
	wn.db.Create(&wn.node)
}

//修改
func (wn *WorkNode) Update() {
}

//删除
func (wn *WorkNode) Delete(conditions map[string]interface{}) {
	query, where := getQueryWhere(conditions)
	wn.db.Delete(query,where)
}


