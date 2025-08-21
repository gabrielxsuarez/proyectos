package config

// Config represents the application configuration loaded from config.yaml
type Config struct {
	Databases map[string]string `yaml:",inline"`
}

// DatabaseConnection represents a database connection configuration
type DatabaseConnection struct {
	Name             string
	ConnectionString string
}

// GetConnections converts the map of databases to a slice of DatabaseConnection
func (c *Config) GetConnections() []DatabaseConnection {
	connections := make([]DatabaseConnection, 0, len(c.Databases))
	for name, connStr := range c.Databases {
		connections = append(connections, DatabaseConnection{
			Name:             name,
			ConnectionString: connStr,
		})
	}
	return connections
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if len(c.Databases) == 0 {
		return ErrNoDatabasesConfigured
	}
	
	for name, connStr := range c.Databases {
		if name == "" {
			return ErrEmptyDatabaseName
		}
		if connStr == "" {
			return ErrEmptyConnectionString
		}
	}
	
	return nil
}