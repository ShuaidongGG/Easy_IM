package main

import (
	"Easy_IM/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/easy_im?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// db.AutoMigrate(&models.UserBasic{})
	// db.AutoMigrate(&models.Contact{})
	// db.AutoMigrate(&models.GroupBasic{})
	// db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Community{})
	// Create

	usr := &models.UserBasic{}

	usr.Name = "gsd1"
	usr.PassWord = "123456"
	usr.LoginTime = time.Now()
	usr.LoginOutTime = time.Now()
	usr.HeartbeatTime = time.Now()
	db.Create(usr)

	// Read
	// var product Product
	// fmt.Println(db.First(usr, 1))
	// db.First(usr, 1)                 // find product with integer primary key
	// db.First(usr, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	// db.Model(usr).Update("Password", "12346")
	// Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(&product, 1)
}
