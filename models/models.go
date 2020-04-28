package models

import (
	"simple-crud/db"
	"time"
)

type Users struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" gorm:"type:varchar(128);not null;index"`
	Age       int       `json:"age" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
}

//func (u *Users) Save() error {
//	return db.DB.Save(u).Error
//}

func CreateTable() {
	//fmt.Printf("initiating db, all exists data will be deleted!\n")
	//fmt.Print("y/n: ")
	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')
	//if text != "y\n" {
	//	log.Println("cancel!")
	//	return
	//}
	db.DB.DropTableIfExists(&Users{})
	_db := db.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	_db.CreateTable(&Users{})

}
