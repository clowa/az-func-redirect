package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clowa/az-func-redirect/src/handlers"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/health", handlers.HealthRouter)
	http.HandleFunc("/healthz", handlers.HealthRouter)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://qunisday.qunis.de", http.StatusMovedPermanently)
	})

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
