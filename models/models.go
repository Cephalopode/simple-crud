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
	mockUsers := []Users{{Name: "Jean-Serge", Age: 45}, {Name: "Gertrude", Age: 79}, {Name: "Octave", Age: 80},
		{Name: "Simone", Age: 45}, {Name: "Henriette", Age: 76}, {Name: "Paul-Maurice", Age: 46}, {Name: "Hectore", Age: 43}, {Name: "Régimbald", Age: 45}}
	for _, item := range mockUsers {
		_db.Save(&item)
	}

}
