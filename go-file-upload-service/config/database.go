package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Database holds the gorm.DB instance
*/
type Database struct {
	DB *gorm.DB
}

/*
NewDatabase function is uses to create new db instance
*/
func NewDatabase(dsn string) *Database {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("%#v\n", db)

	return &Database{
		DB: db,
	}
}

func (db Database) getInstance() *gorm.DB {
	return db.DB
}
