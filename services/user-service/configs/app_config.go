package configs

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var AppConfig Config

type Config struct {
	Server           ServerConfig           `yaml:"server"`
	JWT              JWTConfig              `yaml:"jwt"`
	Database         DatabaseConfig         `yaml:"database"`
	Redis            RedisConfig            `yaml:"redis"`
	Password         PasswordConfig         `yaml:"password"`
	InternalSecurity InternalSecurityConfig `yaml:"internal-security"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type JWTConfig struct {
	Secret             string `yaml:"secret"`
	Expiry             string `yaml:"expiry"`
	RefreshTokenExpiry string `yaml:"refresh-token-expiry"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	DSN      string `yaml:"dsn"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type PasswordConfig struct {
	PasswordResetURL string `yaml:"PasswordResetURL"`
}

type InternalSecurityConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
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
