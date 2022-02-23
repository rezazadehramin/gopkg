package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySQL struct {
	Config Configurator
}

// Connect implements the connector interface for mysql.
func (m *mySQL) Connect() (*gorm.DB, error) {
	db, _ := gorm.Open(mysql.Open(m.Config.DSN()), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.Ping()
	retryCount := 10

	for {
		err := sqlDB.Ping()
		if err != nil {
			if retryCount == 0 {
				return nil, fmt.Errorf("not able to establish connection to database")
			}

			fmt.Sprintf("Could not connect to database. Wait 1 second. %d retries left...", retryCount)
			retryCount--
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	return db, nil
}
