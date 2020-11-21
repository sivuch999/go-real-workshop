package models

import (
	"go-test/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Connect() (err error) {
	dsn := "host=[HOST] user=postgres password=[PASSWORD] dbname=[DB-NAME] port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	inits := []entities.Inits{}
	if err = GormDB.Debug().Raw("SELECT 1+1 AS result").Scan(&inits).Error; err != nil {
		return err
	}
	return nil
}
