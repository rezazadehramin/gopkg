package cache

import (
	"fmt"
	"os"
	"time"
)

// Configurator describes the methods required for a config
// object to be accepted by an engine.
type Configurator interface {
	Addr() string
	Pwd() string
}

// Config holds the default required values to open a connection with
// a remote cache.
type Config struct {
	Host     string
	Port     string
	Password string
}

// Addr returns the connection target tcp/http address.
func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// Pwd returns the authentication password or an empty
// string if no password is defined.
func (c *Config) Pwd() string {
	return c.Password
}

// Engine describes methods exposed by cache engines.
type Engine interface {
	Get(k string) (interface{}, error)
	GetSafe(k string) interface{}
	Set(k string, v interface{}, ttl time.Duration) error
}

// NewENVConfig constructs a configuration object with
// the values found on the environment.
func NewENVConfig() Configurator {
	return &Config{
		Host:     os.Getenv("CACHE_HOST"),
		Port:     os.Getenv("CACHE_PORT"),
		Password: os.Getenv("CACHE_PASSWORD"),
	}
}
