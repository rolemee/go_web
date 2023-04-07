package api

import (
	"time"
	// pq "github.com/lib/pq"
)
const SUCCESS = 10000
const PERMISSIONS_DENIED = 401
const NOT_FOUND = 404
const PARAMETER_MISSING = 9992
const PARAMETER_ERROR = 9995
const SERVER_ERROR = 9994
const WEBSOCKET_RECEIVE_FORMAT_ERROR = "错误的输入"
var WhiteList = []string{"nfts","users"}

type myjson map[string]interface{}

type Response struct{
	Code int `json:"id"`
	Data myjson `json:"data"`
	Message string `json:"message"`
}

type WebsocketAuctionReceive struct{
	NftId int `json:"nft_id"`
	Price float32 `json:"price"`
}
type WebsocketAuctionPayload struct{
	Username string
	AddPrice float32
	TotalPrice float32
}
type User struct{
	Id int `gorm:"primaryKey"`
	NftId int
	UserId int
	Time time.Time

	// Follow_users pq.StringArray `gorm:"type:text[]"`
}