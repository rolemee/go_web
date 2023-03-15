package test

import (
	"outsourcing/module"

	"github.com/gin-gonic/gin"
)
type Response struct{
	Code int `json:"id"`
	Data DataStruct `json:"data"`
	Message string `json:"message"`
}
type DataStruct struct{
	Data []interface{} `json:"data"`
}

func Test(c *gin.Context){

	keyword :=c.Query("keyword")
	index := c.DefaultQuery("index","nfts")
	if keyword == "" || index == ""{
		c.JSON(200, gin.H{
			"code": 9992,
			"data": gin.H{},
			"message": "参数缺失",
		})
	}
	c.JSON(200, gin.H{
		"code": 10000,
		"data": gin.H{"data":module.Search(keyword,index)},
		"message": "查询成功",
	})
	
}