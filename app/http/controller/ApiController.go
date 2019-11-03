package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {

}

type result struct {
	Code    int    		`json:"code"`
	Message string 		`json:"message"`
	Data 	interface{} `json:"data"`
}


/**
	主要用这个做返回
*/
func (a *Api) printSuccessStandard(c *gin.Context, data interface{}) {
	if data == nil {
		data = make(map[string]string)
	}
	result := result{
		Code:    200,
		Message: "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, result)
	return
}
