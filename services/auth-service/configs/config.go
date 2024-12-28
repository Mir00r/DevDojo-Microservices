package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

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
	Secret string `yaml:"secret"`
	Expiry string `yaml:"expiry"`
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

var AppConfig Config

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}
}

func TokenExpiry() time.Duration {
	// Parse the "24h" string into a time.Duration
	duration, err := time.ParseDuration(AppConfig.JWT.Expiry)
	if err != nil {
		log.Fatalf("Failed to parse JWT expiry duration: %v", err)
	}
	return duration
}
