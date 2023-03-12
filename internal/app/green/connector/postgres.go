package connector

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewGormDB() *gorm.DB {
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASS")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, pass, dbname, port)

	postgresConfig := &postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}

	gormConfig := gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	db, err := gorm.Open(postgres.Open(postgresConfig.DSN), &gormConfig)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
