package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"simple-crud/db"
	"time"
)

type Users struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	Name      string    `gorm:"type:varchar(128);not null;index"`
	Age       int       `gorm:"index"`
	CreatedAt time.Time `gorm:"index"`
}

//func (u *Users) Save() error {
//	return db.DB.Save(u).Error
//}

func CreateTable() {
	fmt.Printf("initiating db, all exists data will be deleted!\n")
	fmt.Print("y/n: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if text != "y\n" {
		log.Println("cancel!")
		return
	}
	db.DB.DropTableIfExists(&Users{})
	_db := db.DB.LogMode(true).Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	_db.CreateTable(&Users{})

}
