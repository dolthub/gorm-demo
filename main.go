package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Letter struct {
	Id      int `gorm:"primaryKey"`
	Letter  string
	Date    time.Time
	Lang    string
	Country string
	Region  string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/letters?interpolateParams=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	getOneSchemaDiff(db)
}
