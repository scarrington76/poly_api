package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

type timeHandler struct {
	format string
}

func GETHandler(w http.ResponseWriter, r *http.Request) {

	th := timeHandler{format: time.RFC1123}
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// We use our custom CORS Middleware
	r.Use(CORS)

	fmt.Print("API initiated")
	r.HandleFunc("/tests", GETHandler)
	// r.HandleFunc("/api/insert", POSTHandler)
	// r.HandleFunc("/api/portfolio", GetPortfolio)
	// r.HandleFunc("/api/account", GETAccountValue)
	// r.HandleFunc("/api/indexes", GetIndexes)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
