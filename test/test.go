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

}