package configs

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var AppConfig Config

type Config struct {
	Server           ServerConfig
	JWT              JWTConfig
	Database         DatabaseConfig
	Redis            RedisConfig
	Password         PasswordConfig
	InternalSecurity InternalSecurityConfig
}

type ServerConfig struct {
	Port string
}

type JWTConfig struct {
	Secret             string
	Expiry             string
	RefreshTokenExpiry string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	DSN      string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type PasswordConfig struct {
	PasswordResetURL string
}

type InternalSecurityConfig struct {
	UserName string
	Password string
}

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return err
	}
	log.Println("Configuration loaded successfully")
	return nil
}
