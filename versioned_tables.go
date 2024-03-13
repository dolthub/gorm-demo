package main

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type DoltLog struct {
	CommitHash string `gorm:"column:commit_hash;primaryKey" json:"commit_hash"`
	Committer  string `gorm:"column:committer json:"committer"`
	Email      string `gorm:"column:email json:"email"`
	Date       string `gorm:"column:date json:"date"`
	Message    string `gorm:"column:message json:"message"`
}

func (l *DoltLog) TableName() string {
	return "dolt_log"
}

func getLastCommit(db *gorm.DB) {
	var log DoltLog
	db.First(&log)
}

type SchemaDiff struct {
	from_table_name       string `gorm:"column:from_table_name" json:"from_table_name"`
	to_table_name         string `gorm:"column:to_table_name" json:"to_table_name"`
	from_create_statement string `gorm:"column:from_create_statement" json:"from_create_statement"`
	to_create_statement   string `gorm:"column:to_create_statement" json:"to_create_statement"`
}

func getOneSchemaDiff(db *gorm.DB) {
	var result []SchemaDiff
	db.Raw(
		/* sql: */ "SELECT * FROM dolt_schema_diff(@from_commit, @to_commit)",
		/* values: */ sql.Named("from_commit", "main"), sql.Named("to_commit", "feature"),
	).Scan(&result)
}

type TableDiff struct {
	fromId     int
	fromLetter string
	fromLang   string
	toId       int
	toLetter   string
	toLang     string

	fromCommit     string
	fromCommitDate time.Time
	toCommit       string
	toCommitDate   time.Time
	diffType       string
}

func getOneDiff(db *gorm.DB) {
	var result map[string]interface{}
	db.Raw(
		/* sql: */ "Select * from dolt_diff where to_commit=@to_commit and from_commit=@from_commit",
		/* values: */ sql.Named("from_commit", "main"), sql.Named("to_commit", "feature"),
	).First(&result)
	for k, v := range result {
		fmt.Printf("%s: %#v\n", k, v)
	}

	tableDiff := TableDiff{
		fromId:         result["from_id"].(int),
		fromLetter:     result["from_letter"].(string),
		fromLang:       result["from_lang"].(string),
		toId:           result["to_id"].(int),
		toLetter:       result["to_letter"].(string),
		toLang:         result["to_lang"].(string),
		fromCommit:     result["from_commit"].(string),
		fromCommitDate: result["from_commit_date"].(time.Time),
		toCommit:       result["to_commit"].(string),
		toCommitDate:   result["to_commit_date"].(time.Time),
		diffType:       result["diff_type"].(string),
	}
	fmt.Println(tableDiff)
}
