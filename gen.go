package main

import (
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func generateModels(gormdb *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldWithIndexTag: true,
	})

	g.UseDB(gormdb) // reuse your gorm db

	gormdb.Exec("use 'letters/feature'")
	if gormdb.Error != nil {
		log.Fatal(gormdb.Error)
	}

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
