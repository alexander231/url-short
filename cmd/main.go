package main

import (
	"fmt"
	"github.com/alexander231/url-short/storage"
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
	if err := godotenv.Load("../.env"); err != nil {
		return errors.Wrap(err, "Loading of .env file failed")
	}

	addr := os.Getenv("ADDR")
	pass := os.Getenv("PASS")
	dbStr := os.Getenv("DB")

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return errors.Wrap(err, "Converting DB env var failed")
	}
	rdb := storage.NewRedisStorage(addr, pass, db)
	return nil
}
