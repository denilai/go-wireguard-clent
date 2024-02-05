package main

import (
	"gowg/storage/sqlite"
	"log/slog"
	"os"
)

type Config struct {
}

type User struct {
	id     int
	name   string
	config Config
}

type Environment int

const (
	DEV Environment = iota
	TEST
	PROD
)

func main() {
	env := DEV
	log := SetupLogger(env)
	storage, err := sqlite.New("/home/denilai/repos/gowg/gowg.db")
	if err != nil {
		log.Error("failed to init storage", err)
	}
	_ = storage

}

func SetupLogger(env Environment) *slog.Logger {
	var log *slog.Logger

	switch env {
	case DEV:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case TEST:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case PROD:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
