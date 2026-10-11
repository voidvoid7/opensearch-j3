package main

import (
	"log/slog"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	// JSON to stdout -> Docker journald driver -> journal MESSAGE field
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log = log.With("service", "demo-api")

	for i := 1; ; i++ {
		latency := rand.IntN(500)
		attrs := []any{"request_id", i, "latency_ms", latency, "path", "/api/users"}

		switch r := rand.IntN(10); {
		case r < 2:
			log.Debug("cache lookup", attrs...)
		case r < 7:
			log.Info("handled request", attrs...)
		case r < 9:
			log.Warn("slow request", attrs...)
		default:
			log.Error("request failed", append(attrs, "err", "upstream timeout")...)
		}
		time.Sleep(time.Second)
	}
}