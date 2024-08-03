package config

import (
	"os"

	"github.com/NatalNW7/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("creating database file")
		err = os.MkdirAll("./db", os.ModePerm) 
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQLite error: %v", err)
		return nil, err
	}
	
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error to automigrate schema: %v", err)
		return nil, err
	}

	return db, nil
}