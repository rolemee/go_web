package api

import (
	// "fmt"
	"outsourcing/module"
	"github.com/Lofanmi/pinyin-golang/pinyin"
	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context){
	index := c.DefaultQuery("index","nfts")
	dict := pinyin.NewDict()
	json := make(map[string]interface{}) 
	if !module.In(index,WhiteList){
		c.JSON(200, Response{PARAMETER_ERROR, myjson{} , "参数不合法"})
		return 
	}
	c.BindJSON(&json)
	_, ok1 := json["id"]
	_, ok2 := json["name"]
	_, ok3 := json["describe"]
	if index == "nfts"{
		if !ok3{
			c.JSON(200, Response{PARAMETER_MISSING, myjson{} , "参数缺失"})
			return 
		}
		json["describe_py"] = dict.Convert(json["describe"].(string)," ").None()
	}else if index == "users"{
		if ok3{
			c.JSON(200, Response{PARAMETER_ERROR, myjson{} , "参数不合法"})
			return 
		}
		ok3 = true
	}
	if index == "" ||  len(json) == 0 || !ok1 || !ok2 || !ok3 {
		c.JSON(200, Response{PARAMETER_MISSING, myjson{} , "参数缺失"})
		return 
	}

	json["name_py"] = dict.Convert(json["name"].(string)," ").None()

	res := module.Insert(json,index)
	if res == "enqueued"{
		c.JSON(200, Response{SUCCESS, myjson{}, "添加成功"})
	}else{
		c.JSON(200, Response{SERVER_ERROR, myjson{} , "服务器错误"})
	}
}
