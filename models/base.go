package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xybor/xychat/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// IntializeDB reads all credentials in environments variable and creates a
// connection to the DB.  It also creates a logs directory and a log file to
// save all errors in the application.
func InitializeDB() {
	var err error

	postgres_host := helpers.MustReadEnv("postgres_host")
	postgres_user := helpers.MustReadEnv("postgres_user")
	postgres_dbname := helpers.MustReadEnv("postgres_dbname")
	postgres_port := helpers.MustReadEnv("postgres_port")
	postgres_password := helpers.MustReadEnv("postgres_password")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s",
		postgres_host,
		postgres_user,
		postgres_dbname,
		postgres_port,
		postgres_password,
	)

	_, err = os.Stat("logs")
	if os.IsNotExist(err) {
		os.Mkdir("logs", 0600)
	}

	out, err := os.OpenFile(
		"logs/db.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0600,
	)

	if err != nil {
		log.Panic(err)
	}

	newLogger := logger.New(
		log.New(out, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: newLogger,
		},
	)
	if err != nil {
		log.Panic(err)
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Panic(err)
	}

	err = sqldb.Ping()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("[Xychat] Connecting to database success")
}

// Get the current db struct
func GetDB() *gorm.DB {
	return db
}

// CreateTables will AutoMigrate all tables in application.  If drop_if_exists
// is set to true, it will drop all tables before creating.
func CreateTables(drop_if_exists bool) {
	if drop_if_exists {
		err := db.Migrator().DropTable(
			&User{},
			&Room{},
			&DetailedRoom{},
			&Chat{},
		)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("[Xychat] Dropped all tables in database")
	}

	err := db.AutoMigrate(
		&User{},
		&Room{},
		&DetailedRoom{},
		&Chat{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("[Xychat] Successfully auto-migrate database")
}
