package config

import (
	"os"
	"strconv"
)

type App struct {
	Name              string
	Env               string
	Host              string
	ImageName         string
	AppVersion        string
	JWTExpirationTime string
	JWTSecret         string
	AESEncryptKey     string
}

func GetApp() App {
	return App{
		Name:              getEnv("APP_NAME", "Golang"),
		Env:               getEnv("APP_ENV", "local"),
		Host:              getEnv("APP_HOST", "localhost"),
		ImageName:         getEnv("IMAGE_NAME", ""),
		AppVersion:        getEnv("APP_VERSION", "v0.0.0"),
		JWTExpirationTime: getEnv("JWT_EXPIRATION_TIME", "3600"),
		JWTSecret:         getEnv("JWT_SECRET", ""),
		AESEncryptKey:     getEnv("AES_ENCRYPT_KEY", ""),
	}
}

func GetCurrentEnvironment() string {
	return GetApp().Env
}

type Cache struct {
	Host     string
	Port     string
	Database int
	Password string
	Lifetime string
}

func GetCache() Cache {
	database, _ := strconv.Atoi(getEnv("REDIS_DATABASE", "0"))
	return Cache{
		Host:     getEnv("REDIS_HOST", "127.0.0.1"),
		Port:     getEnv("REDIS_PORT", "6379"),
		Database: database,
		Password: getEnv("REDIS_PASSWORD", ""),
		Lifetime: getEnv("REDIS_LIFETIME", "30s"),
	}
}

type Database struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Timezone string
}

func GetDatabase() Database {
	return Database{
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "5432"),
		Database: getEnv("DB_DATABASE", "golang"),
		Username: getEnv("DB_USERNAME", "root"),
		Password: getEnv("DB_PASSWORD", ""),
		Timezone: getEnv("DB_TIMEZONE", "UTC"),
	}
}

type AWS struct {
	AccessKeyID     string
	SecretAccessKey string
	DefaultRegion   string
	Bucket          string
}

func GetAWS() AWS {
	return AWS{
		AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
		SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
		DefaultRegion:   getEnv("AWS_DEFAULT_REGION", "us-east-1"),
		Bucket:          getEnv("AWS_BUCKET", ""),
	}
}

type Web struct {
	WebUrl string
}

func GetWeb() Web {
	return Web{
		WebUrl: getEnv("WEB_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetString(key, fallback string) string {
	return getEnv(key, fallback)
}
