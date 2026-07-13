package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host      string
	Port      int
	User      string
	Password  string
	Name      string
	EnableSSL bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int64
	JwtSecretKey string
	DB           DBConfig
}

var configurations *Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP_PORT is required")
		os.Exit(1)
	}
	httpPortInt, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("HTTP_PORT must be a valid integer")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT_SECRET_KEY is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB_HOST is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB_PORT is required")
		os.Exit(1)
	}
	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		fmt.Println("DB_PORT must be a valid integer")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB_NAME is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB_USER is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB_PASSWORD is required")
		os.Exit(1)
	}

	enableSSL := os.Getenv("EnableSSL")
	if enableSSL == "" {
		fmt.Println("EnableSSL is required")
		os.Exit(1)
	}
	enableSSLBool, err := strconv.ParseBool(enableSSL)
	if err != nil {
		fmt.Println("EnableSSL must be a valid boolean")
		os.Exit(1)
	}
	dbConfig := DBConfig{
		Host:      dbHost,
		Port:      dbPortInt,
		Name:      dbName,
		User:      dbUser,
		Password:  dbPassword,
		EnableSSL: enableSSLBool,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     httpPortInt,
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}

}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
