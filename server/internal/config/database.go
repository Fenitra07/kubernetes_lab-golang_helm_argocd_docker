package config

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// DatabaseConnex : Fonction pour établir la connexion GORM à la base de données
func DatabaseConnex() *gorm.DB {
	once.Do(func() {
		dbDSN := AppConfiguration.DB_DSN
		if dbDSN == "" {
			log.Fatalln("Empty database configuration")
		}

		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,         // Don't include params in the SQL log
				Colorful:                  false,         // Disable color
			},
		)

		var err error
		// Ouvrir la connexion à la base de données avec GORM
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})

	return db
}
