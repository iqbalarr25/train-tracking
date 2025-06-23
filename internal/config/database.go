package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ctx context.Context
var db *gorm.DB
var sqliteDB *gorm.DB
var cache *redis.Client

func InitDatabase() {
	var err error
	conf := GetDatabase()
	appConf := GetApp()

	logType := logger.Error
	if appConf.Env == "local" {
		logType = logger.Info
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Host, conf.Username, conf.Password, conf.Database, conf.Port, conf.Timezone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logType),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetDBConnection() *gorm.DB {
	return db
}

func InitSqlLiteDatabase() {
	var err error
	sqliteDB, err = gorm.Open(sqlite.Open("./simulation.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
}

func CreateSqlLiteDBConnection() *gorm.DB {
	return sqliteDB
}

func CheckDatabaseHealth() error {
	database, err := db.DB()
	if err != nil {
		return err
	}
	err = database.Ping()
	if err != nil {
		return err
	}
	return nil
}
