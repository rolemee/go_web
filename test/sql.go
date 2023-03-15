package test

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	pq "github.com/lib/pq"
  )
type User struct{
	Id int `gorm:"primaryKey"`
	Username string 
	Follow_users pq.StringArray `gorm:"type:text[]"`
}

func Connect(){
	dsn := "host=47.99.75.203 user=postgres password=C2gGLYdua3KT8J dbname=postgres port=5432"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	user := User{}
	result := db.First(&user)
	fmt.Println(result.RowsAffected)
	fmt.Println(user.Username)
	fmt.Println(user.Follow_users)
	
}
