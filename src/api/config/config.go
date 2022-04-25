package config

import (
	"os"
	"time"
)

var (
	Test        = false
	LoggerLevel = Getenv("LOG_LEVEL", "debug")

	// Default DB:
	DbUsername        = Getenv("DB_USERNAME", "template")
	DbPassword        = Getenv("DB_PASSWORD", "t3mpl473")
	DbDatabase        = Getenv("DB_DATABASE", "template")
	DbHostname        = Getenv("DB_hostname", "localhost")
	DbPort            = 5432
	DbMaxIdleConn     = 5
	DbMaxOpenConn     = 50
	DbConnMaxLifetime = 5 * time.Minute

	// NewRelic - Key for metrics
	NewRelicKey = Getenv("NEWRELIC_KEY", "")

	// DataDog
	DataDogHost = Getenv("HOST_NODE_IP", "")
)

func Getenv(key string, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
