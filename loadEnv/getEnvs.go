package env

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var onc sync.Once
var tokenLifeTime time.Duration
var dsn string

func init() {
	if godotenv.Load("./.env") != nil {
		log.Println("warn: can't find env file!")
	}
}

func GetDSN() string {
	onc.Do(func() {
		dsn = os.Getenv("DSN")
	})
	return dsn
}

func GetTokenLifeTime() time.Duration {
	onc.Do(func() {
		delay := os.Getenv("tokenLifeTime")
		d, err := strconv.Atoi(delay)
		if err != nil {
			d = 1
		}
		tokenLifeTime = time.Duration(d) * time.Minute
	})
	return tokenLifeTime
}
