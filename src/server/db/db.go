package db

import (
	"fmt"
	"log"
	"time"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	utils "eisandbar/filesync/src/server/utils"
)

func InitDB() {
	db, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
  	db.AutoMigrate(&Room{})
  	db.AutoMigrate(&File{})

	fmt.Println("Postgres DB initialized")
}

type Room struct {
	Id int64 `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time
	Owner string
	OwnerId int64
}

type File struct {
	Id int64 `gorm:"primaryKey"`
	UploadedAt time.Time
	Owner string
	OwnerId int64
	RoomId int64
}
type User struct {
	Id int64 `gorm:"primaryKey"`
	Name string
	RoomId int64
}