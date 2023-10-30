package main

import (
	"net/http"
	"os"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	r := createRoutes()
	http.ListenAndServe(":"+getEnv("PORT", "8080"), r)
}
