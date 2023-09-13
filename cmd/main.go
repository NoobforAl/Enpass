package main

import (
	"flag"
	"os"
	"strconv"

	http "github.com/NoobforAl/Enpass/http/v1"
)

var (
	// -S setup server listener
	listen = flag.String("S", "0.0.0.0:1111", "set your ip:port")

	// -dsn setup database file path
	dsn = flag.String("dsn", "./db.sqlite", "database path")

	// secret key for jwt token
	secret_key = flag.String("sk", "", "set secret key")

	// run debug mode app
	debug = flag.Bool("d", false, "run debug mode")

	// how many life jwt token
	life_time = flag.Uint("lf", 10, "life time")
)

func init() {
	flag.Parse()

	os.Setenv("DSN", *dsn)
	os.Setenv("GIN_MODE", "release")
	os.Setenv("SECRET_KEY", *secret_key)
	os.Setenv("LIFETIME", strconv.Itoa(int(*life_time)))

	if *debug {
		os.Setenv("GIN_MODE", "debug")
		os.Setenv("DEBUG_EN_PASS", "true")
	}
}

func main() {
	app := http.HttpApp()
	if err := app.Run(*listen); err != nil {
		panic(err)
	}
}
