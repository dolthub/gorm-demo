package main

import (
	"context"
	"gorm.io/gorm"
	"time"
)

var ticker = time.NewTicker(time.Minute)

func (l *Letter) AfterSave(tx *gorm.DB) (err error) {
	select {
	case <-ticker.C:
		tx.Exec("call dolt_add('letters')")
		if tx.Error != nil {
			tx.Logger.Warn(context.Background(), tx.Error.Error())
		}
		tx.Exec("call dolt_commit('-m', 'periodic commit')")
		if tx.Error != nil {
			tx.Logger.Warn(context.Background(), tx.Error.Error())
		}
		return nil
	default:
	}
	return nil
}

func testTicker(db *gorm.DB) {
	messages := []string{
		"we are all just walking each other home",
		"the fault, dear Brutus, is not in our stars, but in ourselves",
		"Not that I loved Caesar less, but that I loved Rome more",
	}
	var i int = 2
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			insertLetter(db, i, messages[i%len(messages)])
			i++
		default:
		}
	}
}
