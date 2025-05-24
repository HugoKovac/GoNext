package config

import "os"

type DbConfig struct {
	Host string
	Port string
	User string
	Password string
	DbName string
}

type JwtConfig struct {
	Secret string
}

type Config struct {
	Db DbConfig
	Jwt JwtConfig
}

func LoadConfig() *Config {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret"
	}

	return &Config{
		Db: DbConfig{
			Host: dbHost,
			Port: dbPort,
			User: dbUser,
			Password: dbPassword,
			DbName: dbName,
		},
		Jwt: JwtConfig{
			Secret: jwtSecret,
		},
	}
}