package main

import (
	"fmt"
	"github.com/alexander231/url-short/http"
	"github.com/alexander231/url-short/redis"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := godotenv.Load(".env"); err != nil {
		return errors.Wrap(err, "Loading of .env file")
	}

	addr := os.Getenv("ADDR")
	pass := os.Getenv("PASS")
	dbStr := os.Getenv("DB")

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return errors.Wrap(err, "Converting DB env var")
	}

	rdb := redis.NewStorage(addr, pass, db)
	URLService := redis.NewURLService(rdb)

	srv := http.NewServer(URLService)
	srv.Addr = ":8080"

	if err = srv.ListenAndServe(); err != nil {
		return errors.Wrap(err, "Starting http server")
	}
	return nil
}
