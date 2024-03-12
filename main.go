package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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

var timer = time.NewTimer(time.Minute)

func (l *Letter) AfterSave(tx *gorm.DB) (err error) {
	select {
	case <-timer.C:
		timer.Reset(time.Minute)
		tx.Exec("call dolt_commit('-Am', 'periodic commit')")
		tx.
		return tx.Error
	default:
	}
	return nil
}

func readFirstLetter() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/letters?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Get the first record ordered by primary key
	var letter Letter
	log.Println(db.Session(&gorm.Session{DryRun: true}).First(&letter).Statement.SQL.String())
	// SELECT * FROM letters ORDER BY id LIMIT 1;
}

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	generateModels()
}

func generateModels() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/letters?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
