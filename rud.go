package main

import (
	"gorm.io/gorm"
	"log"
	"time"
)

func readFirstLetter(db *gorm.DB) {
	// Get the first record ordered by primary key
	var letter Letter
	log.Println(db.Session(&gorm.Session{DryRun: true}).First(&letter).Statement.SQL.String())
	// SELECT * FROM letters ORDER BY id LIMIT 1;
}

func updateLetter(db *gorm.DB) {
	// fetch
	var letter Letter
	db.First(&letter)
	// modify
	letter.Country = "Italy"
	letter.Date = time.Now()
	// save
	db.Save(&letter)
}

func chainUpdateLetter(db *gorm.DB) {
	db.Model(&Letter{}).Where("country = ?", "Italy").Update("country", "IT")
}

func insertLetter(db *gorm.DB, id int, msg string) {
	newLetter := &Letter{
		Id:     id,
		Letter: msg,
		Date:   time.Now(),
		Lang:   "english",
	}
	db.Save(newLetter)
	// INSERT INTO `letters` (`letter`,`date`,`lang`,`country`,`region`,`id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `letter`=VALUES(`letter`),`date`=VALUES(`date`),`lang`=VALUES(`lang`),`country`=VALUES(`country`),`region`=VALUES(`region`)"
}
