package driver

import (
	"fmt"

	"github.com/murawakimitsuhiro/go-simple-RESTful-api/models"

	"github.com/jinzhu/gorm"
)

type DB struct {
	DB *gorm.DB
}

var db = &DB{}

func ConnectGorm(host, port, uname, pass, dbname string) (*DB, error) {
	dbSourceStr := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", pass, host, port, dbname)

	var err error
	db.DB, err = gorm.Open("mysql", dbSourceStr)

	db.DB.DB()
	db.DB.AutoMigrate(models.Note{})

	return db, err
}
