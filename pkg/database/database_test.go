package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnector_FromENV(t *testing.T) {
	setEnv()
	defer unsetEnv()

	c := NewConnector(NOp, nil)

	db, err := c.Connect()

	assert.Nil(t, err)
	assert.NotEmpty(t, db)
}

func TestNewConnector_FromConfig(t *testing.T) {
	setEnv()
	defer unsetEnv()

	conf := NewENVConfig()
	c := NewConnector(NOp, conf)

	db, err := c.Connect()

	assert.Nil(t, err)
	assert.NotEmpty(t, db)
}

func TestConfig_DSN(t *testing.T) {
	setEnv()
	defer unsetEnv()

	c := NewENVConfig()

	assert.IsType(t, &Config{}, c)
	assert.Equal(t, "root:superSecretPWD!@tcp(https://testi.ng::27654)/test_db?charset=utf8mb4&parseTime=True&loc=Local", c.DSN())

}

func TestNewConnector_MYSQL(t *testing.T) {
	setEnv()
	defer unsetEnv()

	con := NewConnector(MySQL, nil)

	assert.NotNil(t, con)

}

func TestNewConnector_InvalidEngineL(t *testing.T) {
	setEnv()
	defer unsetEnv()

	con := NewConnector(3, nil)

	assert.Nil(t, con)

}

func setEnv() {
	os.Setenv("DB_CONFIG_HOST", "https://testi.ng")
	os.Setenv("DB_CONFIG_PORT", ":27654")
	os.Setenv("DB_CONFIG_NAME", "test_db")
	os.Setenv("DB_CONFIG_USER", "root")
	os.Setenv("DB_CONFIG_PASSWORD", "superSecretPWD!")
	os.Setenv("DB_CONFIG_ENGINE", "mysql")
	os.Setenv("DB_CONFIG_OPTIONS", "charset=utf8mb4&parseTime=True&loc=Local")
}

func unsetEnv() {
	os.Unsetenv("DB_CONFIG_HOST")
	os.Unsetenv("DB_CONFIG_PORT")
	os.Unsetenv("DB_CONFIG_NAME")
	os.Unsetenv("DB_CONFIG_USER")
	os.Unsetenv("DB_CONFIG_PASSWORD")
	os.Unsetenv("DB_CONFIG_ENGINE")
	os.Unsetenv("DB_CONFIG_OPTIONS")
}
