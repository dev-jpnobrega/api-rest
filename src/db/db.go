package infrastructure

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

// DB - database implement
type DB struct {
	database *gorm.DB
}

// Connect - connect to database and return db instance
func (db *DB) Connect(driver string, connectionString string) (dbInstance *gorm.DB, err error) {
	dbInstance, err = gorm.Open(driver, connectionString)

	database = dbInstance
	db.database = dbInstance

	if err != nil {
		log.Fatal(err)
	}

	return
}

// Get - returns database instance
func (db *DB) Get() *gorm.DB {
	return database
}

// GetModel - returns model method's instance
func (db *DB) GetModel(model interface{}) *gorm.DB {
	return database.Model(model)
}
