package configloader

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var dsn string
var secretKey []byte
var tokenLifeTime time.Duration

func init() {
	if godotenv.Load("./.env") != nil {
		log.Println("warn: can't find env file!")
	}

	// get dsn sqlite
	dsn = os.Getenv("DSN")

	// get dsn sqlite
	secretKey = []byte(os.Getenv("SECRET_KEY"))

	// setup life time
	delay := os.Getenv("LIFETIME")
	d, err := strconv.Atoi(delay)
	if d == 0 && err != nil {
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
