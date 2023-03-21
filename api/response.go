package api

const SUCCESS = 10000
const PERMISSIONS_DENIED = 401
const NOT_FOUND = 404
const PARAMETER_MISSING = 9992
const PARAMETER_ERROR = 9995
const SERVER_ERROR = 9994

var WhiteList = []string{"nfts","users"}

type myjson map[string]interface{}

type Response struct{
	Code int `json:"id"`
	Data myjson `json:"data"`
	Message string `json:"message"`
}


