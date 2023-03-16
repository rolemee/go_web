package api

import (
	"outsourcing/module"
	"github.com/gin-gonic/gin"
)

func All(c *gin.Context){
	keyword :=c.Query("keyword")
	index := c.DefaultQuery("index","nfts")
	if !module.In(index,WhiteList){
		c.JSON(200, Response{PARAMETER_ERROR, myjson{} , "参数不合法"})
		return 
	}
	if keyword == "" || index == ""{
		c.JSON(200, Response{PARAMETER_MISSING, myjson{} , "参数缺失"})
		return 
	}
	c.JSON(200, Response{SUCCESS, myjson{"data":module.Search(keyword,index)}, "查询成功"} )
}