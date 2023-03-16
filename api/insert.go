package api

import (
	"github.com/gin-gonic/gin"
	"outsourcing/module"
	// "fmt"
)

func Insert(c *gin.Context){
	index := c.DefaultQuery("index","nfts")
	json := make(map[string]interface{}) 
	if !module.In(index,WhiteList){
		c.JSON(200, Response{PARAMETER_ERROR, myjson{} , "参数不合法"})
		return 
	}
	c.BindJSON(&json)
	res := module.Insert(json,index)
	_, ok := json["id"]
	if index == "" ||  len(json) == 0 || !ok{
		c.JSON(200, Response{PARAMETER_MISSING, myjson{} , "参数缺失"})
		return 
	}

	if res == "enqueued"{
		c.JSON(200, Response{SUCCESS, myjson{}, "添加成功"})
	}
}
