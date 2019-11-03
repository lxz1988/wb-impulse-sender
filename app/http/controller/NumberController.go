package controller

import (
	"github.com/gin-gonic/gin"
	"wb-impulse-sender/conf"
	"wb-impulse-sender/units"
)

type NumberController struct {
	Api
}

func NewNumber() *NumberController {
	return &NumberController{}
}

func (n *NumberController) GetNumber(c *gin.Context) {
	//获取配置文件信息
	cfg := conf.GetConfigInstance()
	//获取号码
	number, _ := units.GetUnitsInstance(cfg, cfg.Number.RedisKey).Redis.LPop(cfg.Number.ListKey)
	n.printSuccessStandard(c, number)
}
