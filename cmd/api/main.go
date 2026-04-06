package main

import (
	"PuzzleLingua/cmd/internal/httpapi"
	"log"
	"net/http"
)

func main() {
	handler := httpapi.NewHandler()

	log.Println("PuzzleLingua API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
