package middleware

import (
	"log"
	"net/http"
)

func Arekta(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Arekta middleware: Before handling request")
		next.ServeHTTP(w, r)
	})
}
