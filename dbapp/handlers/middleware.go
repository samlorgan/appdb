package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic occurred:", err)

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		endpointInfo := fmt.Sprintf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		timestamp := time.Now().Format(time.RFC3339)
		fmt.Printf("[%s] Endpoint call started : %s \n", timestamp, endpointInfo)
		next.ServeHTTP(w, r)
		elapsed := time.Since(start).Milliseconds()
		endpointInfo = fmt.Sprintf("%s %d ms", endpointInfo, elapsed)
		timestamp = time.Now().Format(time.RFC3339)
		fmt.Printf("[%s] Endpoint call complete : %s \n", timestamp, endpointInfo)
	})
}
