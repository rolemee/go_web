package module

import (
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Auction struct{
	Id int `gorm:"primaryKey"`
	NftId int
	UserId int
	Time time.Time
	StartPrice float32
	MinPrice float32
	TimeConsuming int
	EndPrice float32 
	IsEnd bool
	// Follow_users pq.StringArray `gorm:"type:text[]"`
}

var tmp ConfigInfo
var  pgconf= tmp.Readconfig("postgresql.yaml")
var dsn = "host="+pgconf.Ip+" user="+pgconf.User+" password="+pgconf.Password+" dbname="+pgconf.Dbname+" port="+pgconf.Port
var db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
func (Auction) TableName() string {
	return "auction"
}
func QueryAuction(nftId int) Auction {
	auction := Auction{}
	db.Where("is_end=?",false).Where("nft_id=?",nftId).Find(&auction)
	return auction
}
func Updateuction(nftId int,userId int,endPrice float32)  {
	auction := Auction{}
	db.Model(&auction).Where("nft_id=?",nftId).Updates(map[string]interface{}{"user_id": userId, "end_price": endPrice})
}

