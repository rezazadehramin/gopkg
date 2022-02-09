package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySQL struct {
	Config Configurator
}

// Connect implements the connector interface for mysql.
func (m *mySQL) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(m.Config.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
