package driver

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/murawakimitsuhiro/go-simple-RESTful-api/models"
)

type DBConnectionConfig struct {
	DBHost     string
	DBPort     uint
	DBUser     string
	DBPassword string
	DBName     string
}

type DB struct {
	DB *gorm.DB
}

var db = &DB{}

func ConnectGorm(c DBConnectionConfig) (*DB, error) {
	dbSourceStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)

	var err error
	db.DB, err = gorm.Open("mysql", dbSourceStr)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	db.DB.DB()
	db.DB.AutoMigrate(models.Note{})

	return db, err
}
