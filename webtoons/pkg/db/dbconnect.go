package db

import (
	"log"

	"github.com/webtoons/pkg/config"
	"github.com/webtoons/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := cfg.DBURL
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		log.Fatalf("Error connecting to the database: %v", dbErr)
	}

	db.AutoMigrate(
		&domain.User{},
		&domain.Webtoon{},
	)
	return db, dbErr
}
