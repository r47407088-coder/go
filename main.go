package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "embed"
)

//go:embed page.html
var page string

func main() {
	http.HandleFunc("/", serveHello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server:", err)
	}

	http.HandleFunc("/test", test)
}

func serveHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}
func test (w http.ResponseWriter , r *http.Request){
	
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	fmt.Fprint(w, "hello")
}
