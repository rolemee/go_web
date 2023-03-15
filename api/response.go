package api

const SUCCESS = 10000
const PERMISSIONS_DENIED = 401
const NOT_FOUND = 404
const PARAMETER_MISSING = 9992

type myjson map[string]interface{}

type Response struct{
	Code int `json:"id"`
	Data myjson `json:"data"`
	Message string `json:"message"`
}

