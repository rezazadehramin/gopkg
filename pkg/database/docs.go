// Package database contains bindings to connect to
// database engines.
//
// When using the database package you can either provide
// a Configurator or let the package build it by setting
// the following environment variables:
//
//	```
// - DB_CONFIG_HOST
// - DB_CONFIG_PORT
// - DB_CONFIG_NAME
// - DB_CONFIG_USER
// - DB_CONFIG_PASSWORD
// - DB_CONFIG_ENGINE
// - DB_CONFIG_OPTIONS
//  ```
package database
