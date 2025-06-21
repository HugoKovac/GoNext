package config

import "os"

type EnvConfig struct {
	Mode string
}

type CorsConfig struct {
	AllowOrigins string
}

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
	Env EnvConfig
	Cors CorsConfig
}

func LoadConfig() *Config {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "dev"
	}
	allow_origins := os.Getenv("ALLOW_ORIGINS")

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
		Env: EnvConfig{
			Mode: mode,
		},
		Cors: CorsConfig{
			AllowOrigins: allow_origins,
		},
	}
}