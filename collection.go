package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	dbURL  = "postgres://postgres@postgres/postgres?sslmode=disable"
	dbConn *gorm.DB
)

func OverrideDBURL(url string) {
	dbURL = url
}

func DB() *gorm.DB {
	if dbConn != nil {
		return dbConn
	}

	log.Printf("Open postgres @ (%s)", dbURL)
	var err error
	dbConn, err = gorm.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	log.Printf("migrating")
	err = dbConn.AutoMigrate(&Meal{}, &Message{}, &Place{},
		&PlaceVisit{}, &User{}, &Sugar{}, &Food{}).Error
	if err != nil {
		panic(err)
	}
	log.Printf("migrated")
	return dbConn
}

func ILike(db *gorm.DB, field, like string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s ILIKE ?", field), fmt.Sprintf("%%%s%%", like))
}
