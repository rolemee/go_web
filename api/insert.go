package api

import (
	// "fmt"
	"outsourcing/module"
	"github.com/Lofanmi/pinyin-golang/pinyin"
	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context){
	index := c.DefaultQuery("index","nfts")
	json := make(map[string]interface{}) 
	if !module.In(index,WhiteList){
		c.JSON(200, Response{PARAMETER_ERROR, myjson{} , "参数不合法"})
		return 
	}
	c.BindJSON(&json)
	_, ok1 := json["id"]
	_, ok2 := json["name"]
	if index == "" ||  len(json) == 0 || !ok1 || !ok2{
		c.JSON(200, Response{PARAMETER_MISSING, myjson{} , "参数缺失"})
		return 
	}
	dict := pinyin.NewDict()
	json["name_py"] = dict.Convert(json["name"].(string)," ").None()
	res := module.Insert(json,index)

	if res == "enqueued"{
		c.JSON(200, Response{SUCCESS, myjson{}, "添加成功"})
	}
}
