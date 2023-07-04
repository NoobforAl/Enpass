package env

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var dsn string
var tokenLifeTime time.Duration

func init() {
	if godotenv.Load("./.env") != nil {
		log.Println("warn: can't find env file!")
	}

	// get dsn sqlite
	dsn = os.Getenv("DSN")

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

func GetLifeTime() time.Duration {
	return tokenLifeTime
}
