package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"hello-page-and-return-name/config"
	"hello-page-and-return-name/pkg/transport"
)

var DefaultConfig config.Config

func init() {
	DefaultConfig = config.MakeConfig()
}

func Run(cfg config.Config) {

	http.Handle("/hello", transport.Hello())
	http.Handle("/user", transport.User(cfg))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	log.Fatal(http.ListenAndServe(cfg.HTTPAddr, nil))
}

func main() {
	Run(DefaultConfig)
}
