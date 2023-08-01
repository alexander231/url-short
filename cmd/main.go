package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
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

	return nil
}
