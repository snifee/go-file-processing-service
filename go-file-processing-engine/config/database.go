package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
NewDatabase function is uses to create new db instance
*/
func NewDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("Connected to Postgres DB")

	return db
}
