package api

import (
	// "fmt"
	"outsourcing/module"
	"github.com/gin-gonic/gin"
)


func Delete(c *gin.Context){
	index := c.DefaultQuery("index","nfts")
	id :=c.DefaultQuery("id","-1")
	if id == "-1"{
		c.JSON(200, Response{PARAMETER_MISSING, myjson{} , "参数缺失"})
		return 
	}
	if module.Delete(id,index) == "enqueued"{
		c.JSON(200, Response{SUCCESS, myjson{} , "删除成功"})
	}else{
		c.JSON(200, Response{SERVER_ERROR, myjson{} , "服务器错误"})
	}
	
}