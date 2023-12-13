package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/cj-f/audit-demo/pangea_proxy"
	"github.com/pangeacyber/pangea-go/pangea-sdk/v3/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/v3/service/audit"
)

type AuditEvent struct {
	Message string `json:"message"`
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers:", "Origin, Accept, Referer, User-Agent, Content-Type, X-Auth-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middleware again")
	})
}

func main() {
	mux := http.NewServeMux()

	pangeaToken := os.Getenv("PANGEA_TOKEN")
	if pangeaToken == "" {
		log.Fatal("Failed to start server, expecting PANGEA_TOKEN environment variable")
	}

	auditcli, err := audit.New(
		&pangea.Config{
			Token:  pangeaToken,
			Domain: os.Getenv("PANGEA_DOMAIN"),
		},
		audit.WithCustomSchema(AuditEvent{}),
	)

	if err != nil {
		log.Fatal("failed to create audit client")
	}

	pangea := pangea_proxy.New(auditcli)
	mux.HandleFunc("/audit/search", pangea.Search)
	mux.HandleFunc("/audit/results", pangea.Results)
	mux.HandleFunc("/audit/root", pangea.Root)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	log.Println("Starting server on :4000")
	handler := cors.Handler(mux)
	err = http.ListenAndServe(":4000", handler)
	log.Fatal(err)
}
