package postgresql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const errorMessage = "Database connection failed"

func getConnectionString(options *Options) string {
	return fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=product_service sslmode=disable",
		options.PgUsername,
		options.PgPassword)
}

func NewGormDB(options *Options) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(getConnectionString(options)), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("%s: %v", errorMessage, err)
	}

	return db
}

type Options struct {
	PgUsername string
	PgPassword string
	PgDbUrl    string
}
