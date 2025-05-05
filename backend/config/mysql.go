package config

const (
	// Default MySQL Configurations
	MYSQL_USER     = "user"
	MYSQL_PASSWORD = "1234pass"
	MYSQL_HOST     = "127.0.0.1"
	MYSQL_PORT     = "3306"
	MYSQL_DATABASE = "database"
)

type mySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// WithMySQL create a new config instance with MySQL configurations
func (c Config) WithMySQL() Config {
	return Config{
		ServerPort:      c.ServerPort,
		CorsAllowOrigin: c.CorsAllowOrigin,
		MySQL: mySQLConfig{
			User:     getEnvOrDefault("MYSQL_USER", MYSQL_USER),
			Password: getEnvOrDefault("MYSQL_PASSWORD", MYSQL_PASSWORD),
			Host:     getEnvOrDefault("MYSQL_HOST", MYSQL_HOST),
			Port:     getEnvOrDefault("MYSQL_PORT", MYSQL_PORT),
			Database: getEnvOrDefault("MYSQL_DATABASE", MYSQL_DATABASE),
		},
	}
}
