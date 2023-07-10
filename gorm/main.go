package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	BirthDay time.Time
}

func main() {
	dsn := "root:wawawa@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name:     "Asa",
		Email:    "wwfyde@163.com",
		BirthDay: time.Now(),
	})
	var user User
	result := db.First(&user, "name = ?", "Asa")
	//result := db.Take(&user, 6)
	count := result.RowsAffected /* returns count of records found */
	err = result.Error           /*db.First(&user, "name = ?", "wwfyde")*/
	fmt.Println(user.Name, count)

}
