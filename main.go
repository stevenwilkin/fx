package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/stevenwilkin/fx/currencybeacon"

	_ "github.com/joho/godotenv/autoload"
)

var (
	rates json.RawMessage
	m     sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	w.Write(rates)
}

func main() {
	go func() {
		cb := currencybeacon.NewCurrencyBeaconFromEnv()
		ticker := time.NewTicker(10 * time.Minute)

		for {
			if results, err := cb.GetRates(); err != nil {
				slog.Error(err.Error())
			} else {
				m.Lock()
				rates = results
				m.Unlock()
			}

			<-ticker.C
		}
	}()

	port := "8080"
	if wwwPort := os.Getenv("WWW_PORT"); len(wwwPort) > 0 {
		port = wwwPort
	}

	slog.Info("Starting", slog.String("port", port))
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
