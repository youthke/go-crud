package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/youthke/go-crud/pkg/domain/model"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	fmt.Println(conf.DB.User, conf.DB.Password, conf.DB.IP)
	db, err = gorm.Open("mysql", conf.DB.User+":"+conf.DB.Password+"@tcp("+conf.DB.IP+")/crud?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}

	autoMigration()
}

func GetDB() *gorm.DB{
	return db
}

func autoMigration() {
	db.AutoMigrate(model.Book{})
}

