package transport

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"hello-page-and-return-name/config"
	"hello-page-and-return-name/pkg/db"
	"hello-page-and-return-name/pkg/logger"
	"hello-page-and-return-name/pkg/metrics"
)

func Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			metrics.Request.WithLabelValues("GET").Inc()
			fmt.Fprintf(w, "Hello Page")
		}
	})
}
func User(cfg config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		dt := now.Format("15:04:05 - 02.01.2006")

		resp := r.URL.Query()

		switch r.Method {
		case "POST":
			metrics.Request.WithLabelValues("POST").Inc()
			if cfg.SQLiteFile != "" {
				for _, name := range resp["name"] {
					logger.SQLite(cfg.SQLiteFile, name, dt)
					log.Printf("%s: %s", name, dt)
				}
			} else {
				logger.File(cfg.UserLogFile)
				for _, name := range resp["name"] {
					logger.Log.Printf("%s: %s", name, dt)
					log.Printf("%s: %s", name, dt)
				}
			}
		case "GET":
			metrics.Request.WithLabelValues("GET").Inc()
			if cfg.SQLiteFile != "" && len(resp["name"]) != 0 {
				for _, name := range resp["name"] {
					userData := db.GetUser(cfg.SQLiteFile, name)
					fmt.Fprintf(w, "Name: %s\nData:\n", name)
					for _, str := range userData {
						fmt.Fprintf(w, "* %s\n", str)
					}
				}
			}
		}
	})
}
