package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

// Engine are supported database engines.
type Engine int

// Supported DB engines.
const (
	// NOp is a No Operation db engine.
	NOp Engine = iota
	// By default we support mysql which can also
	// be used for MariaDB and potentially AWS Aurora.
	MySQL
)

// Connector exposes a common interface to perform
// tcp connections to different supported database
// engines.
type Connector interface {
	Connect() (*gorm.DB, error)
}

type nop struct {
	Config Configurator
}

// Connect implements the connector interface for NOp.
func (n *nop) Connect() (*gorm.DB, error) {
	return &gorm.DB{Config: &gorm.Config{}}, nil
}

// Configurator describes the methods required for a config
// object to be accepted by a connector.
type Configurator interface {
	DSN() string
}

// Config holds the default required values to open a connection with
// a database engine.
//
// Options needs to be a valid query string (no ? as suffix).
// Valid: parseTime=true
// Invalid: ?parseTime=true
type Config struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Engine   string
	Options  string
}

// DSN handles the information about a specific database that an
// Open Database Connectivity ( ODBC ) driver needs in order to
// connect to it.
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.Options,
	)
}

// NewConnector builds the connector as specified.
//
// If the configurator is nil then it will attempt to
// parse the values from the environment variables predefined,
// if no value can be parsed then it will provide an empty config
// struct.
func NewConnector(e Engine, config Configurator) Connector {
	var conf Configurator
	if config == nil {
		conf = &Config{
			Host:     os.Getenv("DB_CONFIG_HOST"),
			Port:     os.Getenv("DB_CONFIG_PORT"),
			Name:     os.Getenv("DB_CONFIG_NAME"),
			User:     os.Getenv("DB_CONFIG_USER"),
			Password: os.Getenv("DB_CONFIG_PASSWORD"),
			Engine:   os.Getenv("DB_CONFIG_ENGINE"),
			Options:  os.Getenv("DB_CONFIG_OPTIONS"),
		}
	} else {
		conf = config
	}

	switch e {
	case NOp:
		return &nop{
			Config: conf,
		}
	case MySQL:
		return &mySQL{
			Config: conf,
		}
	}

	return nil
}

// NewENVConfig constructs a configuration object from
// the values found on the environment.
func NewENVConfig() Configurator {
	return &Config{
		Host:     os.Getenv("DB_CONFIG_HOST"),
		Port:     os.Getenv("DB_CONFIG_PORT"),
		Name:     os.Getenv("DB_CONFIG_NAME"),
		User:     os.Getenv("DB_CONFIG_USER"),
		Password: os.Getenv("DB_CONFIG_PASSWORD"),
		Engine:   os.Getenv("DB_CONFIG_ENGINE"),
		Options:  os.Getenv("DB_CONFIG_OPTIONS"),
	}
}
