package infra

import (
	"fmt"

	"github.com/dxckboi/hugeman-exam/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(postgres.Open(DatabaseConnectionString()), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("infra::InitDB() - failed to connect to database")
	}
}

func GetDB() *gorm.DB {
	return db
}

func DatabaseConnectionString() string {
	cfg := config.Get().Database
	format := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	return fmt.Sprintf(format, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DatabaseName)
}
