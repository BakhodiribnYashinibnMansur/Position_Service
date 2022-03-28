package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var PerPageSize = 100

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDB       string
	LogLevel         string

	UserServiceHost string
	UserServicePort int

	ProducerServiceHost string
	ProducerServicePort int

	RPCPort string

	PasscodePool   string
	PasscodeLength int
}

// Load loads environment vars and inflates Config
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "postgres"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "qwerty"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 5002))

	c.ProducerServiceHost = cast.ToString(getOrReturnDefault("PRODUCER_SERVICE_HOST", "localhost"))
	c.ProducerServicePort = cast.ToInt(getOrReturnDefault("PRODUCER_SERVICE_PORT", 5003))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":5004"))

	c.PasscodePool = cast.ToString(getOrReturnDefault("PASSCODE_POOL", "0123456789"))
	c.PasscodeLength = cast.ToInt(getOrReturnDefault("PASSCODE_LENGTH", "6"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
