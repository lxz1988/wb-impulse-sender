package pond

import (
	"wb-impulse-sender/conf"
	"wb-impulse-sender/library"
	"wb-impulse-sender/model"
	"sync"
	"time"
)

const (
	NodeBits uint8 = 10 //节点部分位数
	StepBits uint8 = 12 //序列号部分位数
	NodeMax  int64 = -1 ^ (-1 << NodeBits) //1023
	StepMax  int64 = -1 ^ (-1 << StepBits)
	TimeStampLeft  uint8 = NodeBits + StepBits //时间戳部分向左移动的位数
	NodeLeft	   uint8 = StepBits //节点部分向左移动的位数
)
// timestamp 2019-10-01 00:00:00:00 可以使用到2088-10-01
var epoch int64 = 1569859200000

//产品号码的工作节点类
type WorkNode struct {
	mu 	 sync.RWMutex
	ts   int64 //时间戳部分
	node int64 //节点部分
	sn	 int64 //序列部分
}

func NewWorkNode(cfg *conf.Config, driver string) *WorkNode {
	//获取Ip
	ip := library.GetLibraryInstance().Utils.GetServerIp4()
	wn := model.NewWorkNode(cfg, driver)
	conditions := make(map[string]interface{})
	conditions["server_id"] = ip
	//查询当前ip下的ID
	nodes := wn.Select(conditions)
	var nodeId int64
	if len(nodes) > 0 {
		nodeId = nodes[0].ID
	}else {
		//不存在 则插入，生成新的NodeId
		wn.ServerId = ip
		wn.Insert()
		nodeId = wn.ID
	}
	return &WorkNode{
		ts : 0,
		node : nodeId,
		sn : 0,
	}
}

//发号
func (wn *WorkNode) Generate() int64 {
	wn.mu.Lock()
	defer wn.mu.Unlock()
	//获取当前毫秒级时间戳
	now := time.Now().UnixNano() / 1e6
	//若时间戳一致（说明同一毫秒内需要生成多个UUID）
	if wn.ts == now {
		wn.sn++
		if wn.sn > StepMax { //同一毫秒内最多允许生成4096(0、1、2、3...4095)个序列号
			for now <= wn.ts {
				now = time.Now().UnixNano() / 1e6
			}
		}
	}else{
		wn.sn = 0
	}
	wn.ts = now
	//按位取号码
	result := int64((wn.ts - epoch) << TimeStampLeft | wn.node << NodeLeft | wn.sn)
	return result
}
