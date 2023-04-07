package test
import "testing"
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"time"
	pq "github.com/lib/pq"
  )
type User struct{
	Id int `gorm:"primaryKey"`
	Username string 
	Password string
	Follow_users pq.StringArray `gorm:"type:text[]"`
}
type Auction struct{
	Id int `gorm:"primaryKey"`
	NftId int
	UserId int
	Time time.Time
	StartPrice float32
	MinPrice float32
	TimeConsuming int
	EndPrice float32
	// Follow_users pq.StringArray `gorm:"type:text[]"`
}
func (Auction) TableName() string {
	return "auction"
}

func TestConnect(t *testing.T){
	dsn := "host=47.99.75.203 user=postgres password=C2gGLYdua3KT8J dbname=postgres port=5432"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	auction := Auction{}
	res := db.Where("id=?",1).Find(&auction)
	fmt.Println(res.RowsAffected)
	fmt.Println(auction)
}
