package config

import "errors"

var (
	ErrConfigFileNotFound      = errors.New("config.yaml file not found")
	ErrInvalidConfigFormat     = errors.New("invalid config.yaml format")
	ErrNoDatabasesConfigured   = errors.New("no databases configured in config.yaml")
	ErrEmptyDatabaseName       = errors.New("database name cannot be empty")
	ErrEmptyConnectionString   = errors.New("connection string cannot be empty")
)