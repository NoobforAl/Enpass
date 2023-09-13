package configloader

import (
	"os"
	"strconv"
	"time"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/joho/godotenv"
)

const (
	DSN        = "DSN"
	LIFE_TIME  = "LIFETIME"
	GIN_MODE   = "GIN_MODE"
	SECRET_KEY = "SECRET_KEY"
)

var (
	dsn           string
	gin_mode      string
	secretKey     []byte
	tokenLifeTime time.Duration
)

func EnvInit(log contract.Logger) {
	if godotenv.Load("./.env") != nil {
		log.Warn("warn: can't find env file!")
	}

	// get dsn sqlite
	dsn = os.Getenv(DSN)

	gin_mode = os.Getenv(GIN_MODE)

	// get dsn sqlite
	secretKey = []byte(os.Getenv(SECRET_KEY))

	// setup life time
	delay := os.Getenv(LIFE_TIME)
	d, err := strconv.Atoi(delay)
	if d == 0 || err != nil {
		d = 1
	}
	tokenLifeTime = time.Duration(d) * time.Minute
}

func GetDSN() string {
	return dsn
}

func GetSecretKey() []byte {
	return secretKey
}

func GetLifeTime() time.Duration {
	return tokenLifeTime
}

func GetGinMode() string {
	return gin_mode
}
